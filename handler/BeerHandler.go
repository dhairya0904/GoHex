package handler

import (
	"context"

	BeerService "github.com/dhairya0904/GolangBaseServiceModel/BeerService"
)

type BeerHandler struct {
}

func NewBeerHandler() *BeerHandler {
	return &BeerHandler{}
}

func (beerHandler *BeerHandler) GetBeer(ctx context.Context, beerRequest *BeerService.BeerRequest) (*BeerService.BeerResponse, error) {
	return &BeerService.BeerResponse{Beers: []string{"Tuborg", "Corona"}}, nil
}
