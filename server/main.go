package main

import (
	"context"
	"fmt"
	"net"

	"example.com/grpcdemo/shared"
	"google.golang.org/grpc"
)

func main() {
	// Set address for this server
	fmt.Print("Server address: ")
	var address string
	fmt.Scanln(&address)

	go startServer(address)
	fmt.Scanln()
}

// Implement DemoServiceServer interface
type DemoServiceServer struct {
	shared.UnimplementedDemoServiceServer
	// Server variables goes here
	// timestamp int
	// mu        sync.Mutex
}

func newDemoServiceServer() *DemoServiceServer {
	return &DemoServiceServer{
		// timestamp: 0,
	}
}

func (s *DemoServiceServer) Poke(ctx context.Context, finger *shared.Finger) (*shared.Reaction, error) {
	return &shared.Reaction{
		Insults: []string{
			"Screw you",
			">:/",
			finger.Description + " back at you!",
		},
	}, nil
}

func startServer(address string) {
	fmt.Printf("Starting server at address %s\n\r", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		panic(err.Error())
	}

	var opts []grpc.ServerOption
	// ServerOption configuration goes here

	grpcServer := grpc.NewServer(opts...)

	shared.RegisterDemoServiceServer(grpcServer, newDemoServiceServer())

	fmt.Println("Server starting...")
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err.Error())
	}
}
