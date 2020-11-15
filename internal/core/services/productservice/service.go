package productservice

import (
	"rotteneggs/productservice/internal/core/domain"
	"rotteneggs/productservice/internal/core/ports"
)

type service struct {
	firestoreRepository ports.FirestoreRepository
}

// New : Creates product service
// Parameter: Implementation of firestoreRepository
func New(firestoreRepository ports.FirestoreRepository) *service {
	return &service{
		firestoreRepository: firestoreRepository,
	}
}

func (srv *service) Get(id string) (domain.Product, error) {
	product, err := srv.firestoreRepository.Get(id)
	if err != nil {

		return domain.Product{}, err
	}
	return product, nil
}
