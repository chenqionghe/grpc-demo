package main

import (
	"app/lightweight"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := lightweight.NewGymClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.BodyBuilding(ctx, &lightweight.Person{
		Name:    "chenqionghe",
		Actions: []string{"深蹲", "卧推", "硬拉"},
	})
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("code: %d, msg: %s\n", r.Code, r.Msg)
}
