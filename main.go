package main

import (
	"net"
	"rotteneggs/baseservice/handler"

	"github.com/dhairya0904/GolangBaseServiceModel/BeerService"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	server := grpc.NewServer()
	beerServiceHandler := handler.NewBeerHandler()

	BeerService.RegisterBeerServiceServer(server, beerServiceHandler)
	reflection.Register(server)
	l, _ := net.Listen("tcp", ":9092")
	server.Serve(l)
}
