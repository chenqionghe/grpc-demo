package main

import (
	"app/lightweight"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

// server继承自动生成的服务类
type server struct {
	lightweight.UnimplementedGymServer
}

// 服务端必须实现了相应的接口BodyBuilding
func (s *server) BodyBuilding(ctx context.Context, in *lightweight.Person) (*lightweight.Reply, error) {
	fmt.Printf("%s正在健身, 动作: %s\n", in.Name, in.Actions)
	return &lightweight.Reply{Code: 0, Msg: "ok",}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	lightweight.RegisterGymServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
