package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "github.com/stinkyfingers/modularPlatform/modules/go_example"
	"google.golang.org/grpc"
)

var (
	port = flag.String("p", "9999", "port that module runs on")
)

func main() {
	flag.Parse()
	err := start()
	if err != nil {
		log.Fatal(err)
	}
}

func start() error {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", *port), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewPlatformClient(conn)
	message, err := client.RegisterModule(context.Background(), &pb.Module{Name: "test_go_module", Port: "9999"})
	if err != nil {
		return err
	}
	fmt.Println("MESSAGE: ", message)
	return nil
}
