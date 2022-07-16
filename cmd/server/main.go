package main

import (
	"fmt"
	"pokerpc/server"
)

func main() {
	fmt.Println("Starting...")

	// server := server.NewServer(server.Config{})
	server := server.NewServer(server.Config{
		Protocol: "tcp",
		Host:     "127.0.0.1",
		Port:     "3333",
	})

	err := server.Serve()
	if err != nil {
		panic(err)
	}
}
