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
	addr := fmt.Sprintf("%s:%s", s.config.HostName, s.config.Port)
	listener, err := net.Listen(s.config.Protocol, addr)
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	proto.RegisterPokemonServiceServer(srv, s)

	if err := srv.Serve(listener); err != nil {
		return err
	}
	return nil
}

func (s server) GetList(c context.Context, req *proto.ListRequest) (list *proto.PokemonList, err error) {
	result := s.fetch(req.Limit, req.Offset)
	pkmList := make([]*proto.Pokemon, 0)
	list = &proto.PokemonList{}
	for _, v := range result.Results {
		pkmList = append(pkmList, &proto.Pokemon{
			Id:        0,
			Name:      v.Name,
			IsDefault: false,
			BaseXp:    10,
		})
	}

	list.Pokemon = pkmList
	return
}

func (s server) fetch(limit, offset int32) PokemonPage {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon")
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
