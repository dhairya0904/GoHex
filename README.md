# GolangBaseService
A generic framework to write microservice

All proto files installed at one place: https://github.com/dhairya0904/GolangBaseServiceModel

## Api Flow

handler -> component(Business logic) -> adapter(Optional) -> external(All the external service calls)


# Dependencies
* Install Protobuf compiler
* Install Golang
* Make sure GOPATH is set properly

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