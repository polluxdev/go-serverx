package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/polluxdev/go-serverx/example/grpc/model"
	grpcclient "github.com/polluxdev/go-serverx/grpc/client"
)

func main() {
	// New connection
	conn, err := grpcclient.New(grpcclient.Target("localhost:50051"))
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}
	defer conn.Close()

	// New service
	service := model.NewUserServiceClient(conn.Conn())

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Send request to server
	response, err := service.Create(ctx, &model.CreateUserRequest{Name: "bob", Email: "bob@example.com"})
	if err != nil {
		log.Fatalf("failed to create: %v\n", err)
	}

	// Print response
	fmt.Printf("response: %s\n", response.GetId())
}
