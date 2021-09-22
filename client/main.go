package main

import (
	"context"
	"fmt"
	mygrpc "grpc-demo/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := mygrpc.NewRestaurantLikeSeviceClient(cc)
	request := &mygrpc.RestaurantLikeStatRequest{ResIds: []int32{1, 2, 3}}

	resp, _ := client.GetRestaurantLikeStat(context.Background(), request)
	fmt.Printf("Recieved response => [%v]\n", resp.Result)
}