package productservice

import (
	"context"
	"rotteneggs/productservice/internal/core/domain"
	constants "rotteneggs/productservice/pkg/constants"
	"rotteneggs/productservice/pkg/exceptions"
	"testing"

	"github.com/rotten-eggs/ServiceModel/ProductService"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/pkg/errors"
)

type mockProductService struct{}

var mockProductServiceGetCall func(id string) (domain.Product, error)

func (mockProductService *mockProductService) Get(id string) (domain.Product, error) {
	return mockProductServiceGetCall(id)
}

func TestGet(t *testing.T) {
	mockProductServiceGetCall = func(id string) (domain.Product, error) {
		return constants.ProductDetailsStruct, nil
	}

	productServiceHandler := New(&mockProductService{})
	input := ProductService.GetProductDetailsInput{ProductId: "testId"}
	_, err := productServiceHandler.GetProductDetails(context.TODO(), &input)

	if err != nil {
		t.Errorf(constants.UnexpectedError)
	}
}

func TestGetDependencyFailure(t *testing.T) {
	mockProductServiceGetCall = func(id string) (domain.Product, error) {
		return domain.Product{}, errors.Wrap(&exceptions.DependencyFailureException{errors.New("exception")}, "error")
	}

	productServiceHandler := New(&mockProductService{})
	input := ProductService.GetProductDetailsInput{ProductId: "testId"}
	_, err := productServiceHandler.GetProductDetails(context.TODO(), &input)

	if err == nil || !assertRPCCodes(err, codes.Internal) {
		t.Errorf(constants.ExceptionShouldBeThrown)
	}
}

func TestGetInvalidArgument(t *testing.T) {
	mockProductServiceGetCall = func(id string) (domain.Product, error) {
		return domain.Product{}, errors.Wrap(&exceptions.InvalidArgumentException{errors.New("exception")}, "error")
	}

	productServiceHandler := New(&mockProductService{})
	input := ProductService.GetProductDetailsInput{ProductId: "testId"}
	_, err := productServiceHandler.GetProductDetails(context.TODO(), &input)

	if err == nil || !assertRPCCodes(err, codes.InvalidArgument) {
		t.Errorf(constants.ExceptionShouldBeThrown)
	}
}

func TestGetUnknwonException(t *testing.T) {
	mockProductServiceGetCall = func(id string) (domain.Product, error) {
		return domain.Product{}, errors.Wrap(errors.New("exception"), "error")
	}

	productServiceHandler := New(&mockProductService{})
	input := ProductService.GetProductDetailsInput{ProductId: "testId"}
	_, err := productServiceHandler.GetProductDetails(context.TODO(), &input)

	if err == nil || !assertRPCCodes(err, codes.Unknown) {
		t.Errorf(constants.ExceptionShouldBeThrown)
	}
}

func assertRPCCodes(err error, c codes.Code) bool {
	if e, ok := status.FromError(err); ok {
		if e.Code() != c {
			return false
		}
	} else {
		return false
	}
	return true
}
