package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"pokerpc/config"
	"pokerpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type PokemonItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Pokemon struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	BaseXP    int    `json:"base_experience"`
	IsDefault bool   `json:"is_default"`
	Sprite    string `json:"sprite"`
	Types     string `json:"types"`
}

type PokemonPage struct {
	Count    int           `json:"count"`
	Next     *string       `json:"next"`
	Previous *string       `json:"previous"`
	Results  []PokemonItem `json:"results"`
}

type Server interface {
	Serve() error
}

type server struct {
	config config.Config
	proto.UnimplementedPokemonServiceServer
}

func NewServer(c config.Config) Server {
	return &server{config: c}
}

func (s server) Serve() error {
	addr := fmt.Sprintf(":%s", s.config.Port)
	listener, err := net.Listen(s.config.Protocol, addr)
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	proto.RegisterPokemonServiceServer(srv, s)
	reflection.Register(srv)
	if err := srv.Serve(listener); err != nil {
		return err
	}
	return nil
}

func (s server) GetServerInfo() {

}

func (s server) GetList(c context.Context, req *proto.ListRequest) (list *proto.PokemonList, err error) {
	result := s.fetch(req.Limit, req.Offset)

	pokemonList := <-fetchAll(result)

	pkmList := make([]*proto.Pokemon, 0)
	list = &proto.PokemonList{}

	for _, v := range pokemonList {
		pkmList = append(pkmList, &proto.Pokemon{
			Id:        int32(v.ID),
			Name:      v.Name,
			IsDefault: false,
			BaseXp:    int32(v.BaseXP),
			Sprite:    v.Sprite,
			Types:     v.Types,
		})
	}

	list.Pokemon = pkmList
	return list, err
}

func (s server) Hello(ctx context.Context, in *proto.Empty) (*proto.Greeting, error) {
	return &proto.Greeting{Message: "Hello World x2"}, nil
}

func (s server) fetch(limit, offset int32) PokemonPage {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon?limit=%d&offset=%d", limit, offset)
	resp, err := http.Get(url)
	if err != nil {
		log.Default().Println("Can't fetch from pokeapi.co")
	}

	body, err := ioutil.ReadAll(resp.Body)

	list := new(PokemonPage)
	err = json.Unmarshal(body, list)

	if err != nil {
		log.Default().Println(err)
	}
	return *list
}
