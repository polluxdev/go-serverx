package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"github.com/polluxdev/go-serverx/example/grpc/model"
	grpcserver "github.com/polluxdev/go-serverx/grpc/server"
	"google.golang.org/grpc"
)

type handler struct {
	model.UnimplementedUserServiceServer
}

func (h *handler) Create(_ context.Context, in *model.CreateUserRequest) (*model.CreateUserResponse, error) {
	id := uuid.NewString()
	fmt.Printf("id: %s\n", id)
	fmt.Printf("user: %+v\n", in)
	return &model.CreateUserResponse{Id: id}, nil
}

func main() {
	// New grpc server and registration
	grpcServer := grpc.NewServer()
	model.RegisterUserServiceServer(grpcServer, &handler{})

	// New server
	server, err := grpcserver.New(grpcServer)
	if err != nil {
		log.Fatalf("failed to server: %v", err)
	}

	// Start server
	server.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case <-interrupt:
		log.Print("shutting down")
	case err := <-server.Notify():
		log.Fatal(err)
	}

	// Shutdown
	err = server.Shutdown()
	if err != nil {
		log.Fatal(err)
	}
}
