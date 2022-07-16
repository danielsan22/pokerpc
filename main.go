package main

import (
	"log"
	"pokerpc/config"
	"pokerpc/server"
)

func main() {
	config := *config.NewConfig()
	log.Println(config)
	server := server.NewServer(config)

	err := server.Serve()
	if err != nil {
		panic(err)
	}
}
