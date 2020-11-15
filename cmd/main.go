package main

import (
	"net"
	service "rotteneggs/productservice/internal/core/services/productservice"
	handler "rotteneggs/productservice/internal/handlers/productservice"
	repo "rotteneggs/productservice/internal/repositories/dao"

	"github.com/rotten-eggs/ServiceModel/ProductService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	firestoreRepository := repo.NewProductDataFirestore()
	productService := service.New(firestoreRepository)
	productHanndler := handler.New(productService)

	server := grpc.NewServer()
	ProductService.RegisterProductServiceServer(server, productHanndler)
	reflection.Register(server)
	l, _ := net.Listen("tcp", ":9092")
	server.Serve(l)
}
