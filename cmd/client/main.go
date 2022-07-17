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
	addr := fmt.Sprintf("%s:%s", "gorpc.onrender.com", "50051")

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatalf("impossible connect: %v", err)
	}
	client := proto.NewPokemonServiceClient(conn)
	ctx := context.Background()
	res, err := client.GetList(ctx, &proto.ListRequest{Limit: 20, Offset: 0})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
