package main

import (
	"golang.org/x/net/context"
	et "grpctest/exporttask"
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"time"
)

const (
	port = ":50050"
)

type server struct {}

func (s *server)CreateTask(ctx context.Context,in *et.TaskInput) (*et.TaskOutput, error) {
	return &et.TaskOutput{Worker:"yan", Time:time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006")}, nil
}

func (s *server)DelTask(ctx context.Context, in *et.TaskInput) (*et.TaskOutput, error) {
	return &et.TaskOutput{Worker:"yan", Time:time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006")}, nil
}

func main()  {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	et.RegisterTaskServersServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}