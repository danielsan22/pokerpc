package main

import (
	"context"
	"fmt"
	"log"

	"pokerpc/proto"

	"google.golang.org/grpc"
)

type Client struct {
	cc grpc.ClientConn
}

func main() {
	// addr := fmt.Sprintf("%s:%s", "throbbing-water-7344.fly.dev", "443")
	// addr := fmt.Sprintf("%s:%s", "45.79.33.217", "443")
	addr := fmt.Sprintf("%s:%s", "127.0.0.1", "3333")

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatalf("impossible connect: %v", err)
	}
	client := proto.NewPokemonServiceClient(conn)
	ctx := context.Background()
	res, err := client.GetList(ctx, &proto.ListRequest{Limit: 10, Offset: 10})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
