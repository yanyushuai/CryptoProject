package main

import (
	"google.golang.org/grpc"
	"log"
	et "grpctest/exporttask"
	"golang.org/x/net/context"
	"time"
)

const (
	address 	= "localhost:50050"
)

func main()  {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := et.NewTaskServersClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	o, err :=c.CreateTask(ctx, &et.TaskInput{TaskId:"1", TaskName:"task1"})
	if err != nil {
		log.Fatalf("could not create task : %v", err)
	}

	log.Printf("Worker : %s , Time : %s", o.Worker, o.Time)

	o, err = c.DelTask(ctx, &et.TaskInput{TaskId:"1", TaskName:"task1"})
	if err != nil {
		log.Fatalf("could not create task : %v", err)
	}

	log.Printf("Worker : %s , Time : %s", o.Worker, o.Time)

}