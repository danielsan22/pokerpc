package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type pkmDetail struct {
	index int
	data  Pokemon
}

// for
// go routine
// waitGroup .add
// channel to send result x <- result

func fetchAll(page PokemonPage) chan []Pokemon {
	channelResult := make(chan []Pokemon)
	go func() {
		size := len(page.Results)
		result := make([]Pokemon, size)
		quit := make(chan bool)
		c := make(chan pkmDetail)

		var wg sync.WaitGroup
		wg.Add(size)

		go func() {
			wg.Wait()
			quit <- true
		}()

		for i, p := range page.Results {
			go func(index int, item PokemonItem, cc chan pkmDetail) {
				getDetail(index, item, cc)
				wg.Done()
			}(i, p, c)
		}

		for {
			select {
			case output := <-c:
				result[output.index] = output.data
			case <-quit:
				channelResult <- result
				return
			}
		}
	}()
	return channelResult
}

func getDetail(index int, item PokemonItem, c chan pkmDetail) {
	resp, err := http.Get(item.URL)

	if err != nil {
		log.Default().Println("Can't fetch from pokeapi.co")
	}

	body, err := ioutil.ReadAll(resp.Body)

	info := new(PokemonInfo)
	err = json.Unmarshal(body, info)

	if err != nil {
		log.Default().Println(err)
	}
	r := new(pkmDetail)
	r.index = index
	r.data = Pokemon{
		ID:        info.ID,
		Name:      info.Name,
		BaseXP:    info.BaseExperience,
		IsDefault: info.IsDefault,
		Sprite:    info.Sprites.Other.Home.FrontDefault,
		Types:     "some,types",
	}
	c <- *r
}
