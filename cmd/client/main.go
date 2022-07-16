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
	addr := fmt.Sprintf("%s:%s", "127.0.0.1", "3333")
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("impossible connect: %v", err)
	}
	client := proto.NewPokemonServiceClient(conn)
	ctx := context.Background()
	res, err := client.GetList(ctx, &proto.ListRequest{Limit: 20, Offset: 0})

	fmt.Println(res)
	fmt.Println(err)
}
