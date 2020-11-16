# GoHex
A generic framework(Hexagonal Architecture) to write microservice

ProtoFile: https://github.com/dhairya0904/GolangBaseServiceModel

## How to write code?

https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3

# Dependencies
* Install Protobuf compiler
* Install Golang
* Make sure GOPATH is set properly
* To use prive repo with go module refer: https://medium.com/swlh/go-modules-with-private-git-repository-3940b6835727

# Testing: grpccurl to test grpc calls

## Installation
* https://github.com/fullstorydev/grpcurl

## Usage 

### Get All servers running on port
```
grpcurl --plaintext localhost:9092 list
```

### Get Methods for service
```
grpcurl --plaintext localhost:9092 list BeerService
```

### Describe methods for any service
```
grpcurl --plaintext localhost:9092 describe BeerService.GetBeer
```

### Call method 
```
grpcurl --plaintext -d '{"number": '1'}' localhost:9092 BeerService.GetBeer
```
