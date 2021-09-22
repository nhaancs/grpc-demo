package main

import (
	"context"
	"log"
	"net"

	mygrpc "grpc-demo/pb"

	"google.golang.org/grpc"
)

type server struct {
	mygrpc.UnimplementedRestaurantLikeSeviceServer
}

func (s *server) GetRestaurantLikeStat(
	ctx context.Context,
	req *mygrpc.RestaurantLikeStatRequest,
) (*mygrpc.RestaurantLikeStatResponse, error) {
	return &mygrpc.RestaurantLikeStatResponse{Result: map[int32]int32{1: 1, 2: 4}}, nil
}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln("Failed to listen: ", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to ther server
	mygrpc.RegisterRestaurantLikeSeviceServer(s, &server{})
	// Serve gRPC server
	log.Printf("Serving gRPC on %v ...", address)
	log.Fatal(s.Serve(lis))
}