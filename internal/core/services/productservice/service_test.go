package productservice

import (
	constants "rotteneggs/productservice/common/constants"
	"rotteneggs/productservice/common/exceptions"
	"rotteneggs/productservice/internal/core/domain"
	"testing"

	"github.com/pkg/errors"
)

type mockFirestoreRepository struct {
}

var firestoreRepositoryMockCall func(id string) (domain.Product, error)

func (mockFirestoreRepository *mockFirestoreRepository) Get(id string) (domain.Product, error) {
	return firestoreRepositoryMockCall(id)
}

func TestGet(t *testing.T) {

	firestoreRepositoryMockCall = func(id string) (domain.Product, error) {
		return constants.ProductDetailsStruct, nil
	}

	firestoreRepo := New(&mockFirestoreRepository{})
	actual, err := firestoreRepo.Get("testId")

	if err != nil {
		t.Errorf(constants.UnexpectedError)
	}
	if actual.Creator != constants.Creator {
		t.Errorf("Got wrong creator")
	}
}

func TestGetOnFailure(t *testing.T) {

	firestoreRepositoryMockCall = func(id string) (domain.Product, error) {
		return domain.Product{}, errors.Wrap(&exceptions.DependencyFailureException{errors.New("exception")}, "error")
	}

	firestoreRepo := New(&mockFirestoreRepository{})
	_, err := firestoreRepo.Get("testId")

	if err == nil {
		t.Errorf(constants.ExceptionShouldBeThrown)
	}
}
