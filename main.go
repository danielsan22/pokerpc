package main

import (
	"pokerpc/config"
	"pokerpc/server"
)

func main() {
	config := *config.NewConfig()
	server := server.NewServer(config)

	err := server.Serve()
	if err != nil {
		panic(err)
	}
}
