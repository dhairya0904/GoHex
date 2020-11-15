package dao

import (
	"context"
	"fmt"
	"rotteneggs/productservice/internal/core/domain"
	"rotteneggs/productservice/pkg/exceptions"
	logger "rotteneggs/productservice/pkg/logging"

	"github.com/pkg/errors"

	"github.com/mitchellh/mapstructure"
)

//SqF9Bne0OMNJhbGY3j1X - Test key for database

type productDataFirestore struct {
	Client FirestoreClientInterface
}

// Imterface to interact with firestore
func NewProductDataFirestore() *productDataFirestore {
	return &productDataFirestore{Client: &FirestoreClient{}}
}

// Will be used for mocking purpose in test until we figure out the better way to do so
func GetProductDataFirestore(client FirestoreClientInterface) *productDataFirestore {
	return &productDataFirestore{Client: client}
}

type Product struct {
	Creator     string
	Description string
	SourceUrl   string
	Tags        []string
	Title       string
	Type        string
}

const (
	TimeLimitExceededError string = "Firestore call, Error, timit limit exceeded"
	DataBaseFailureError   string = "Firestore call, Error, call to database failed"
)

func (dataBase *productDataFirestore) Get(id string) (domain.Product, error) {
	log := logger.NewLogger()

	data, err := dataBase.Client.Get(id)

	if err != nil {
		p := domain.Product{}
		switch err {
		case context.DeadlineExceeded:
			return p, errors.Wrap(&exceptions.RecoverableException{err}, TimeLimitExceededError)
		default:
			return p, errors.Wrap(&exceptions.DependencyFailureException{err}, DataBaseFailureError)
		}
	}

	var product Product
	mapstructure.Decode(data, &product)

	log.Info(fmt.Sprintf("Result from database: %+v \n", product))
	return domain.Product{
		ID:          id,
		Name:        product.Title,
		Type:        product.Type,
		Url:         product.SourceUrl,
		Description: product.Description,
		Tags:        product.Tags,
		Creator:     product.Creator,
	}, nil
}
