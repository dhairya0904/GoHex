package ports

import (
	"rotteneggs/productservice/internal/core/domain"
)

// FirestoreRepository - Interface to talk to Firestore database
type FirestoreRepository interface {
	Get(id string) (domain.Product, error)
}

// ProductService - this is where all logic resides
type ProductService interface {
	Get(id string) (domain.Product, error)
}
