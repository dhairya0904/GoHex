package productservice

import (
	"context"
	"fmt"
	"rotteneggs/productservice/internal/core/ports"
	"rotteneggs/productservice/pkg/exceptions"
	logger "rotteneggs/productservice/pkg/logging"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/pkg/errors"
	"github.com/rotten-eggs/ServiceModel/ProductService"
)

type ProductDetailsHandler struct {
	productService ports.ProductService
}

// New - Creates handler for product service
// Paramter - Implementation of {ports.ProductService}
func New(productService ports.ProductService) *ProductDetailsHandler {
	return &ProductDetailsHandler{productService: productService}
}

const internalError = "Internal error while calling GetProductDetails:ProductService"
const unknwonError = "Unknwon error while calling GetProductDetails:ProductService"
const invalidArgumentError = "Invaid argument exception while calling GetProductDetails:ProductService"

// GetProductDetails - Method that fires up when rpc method GetProductDetails if called for productService.
func (handler *ProductDetailsHandler) GetProductDetails(ctx context.Context, request *ProductService.GetProductDetailsInput) (*ProductService.GetProductDetailsResponse, error) {
	log := logger.NewLogger()
	result, err := handler.productService.Get(request.GetProductId())

	if err != nil {

		switch err := errors.Cause(err).(type) {
		case *exceptions.DependencyFailureException:
			log.Error(fmt.Printf("%s: %+v", internalError, err))
			return &ProductService.GetProductDetailsResponse{}, status.Errorf(
				codes.Internal,
				internalError)
		case *exceptions.InvalidArgumentException:
			log.Error(fmt.Printf("%s: %+v", invalidArgumentError, err))
			return &ProductService.GetProductDetailsResponse{}, status.Errorf(
				codes.InvalidArgument,
				invalidArgumentError)
		default:
			log.Error(fmt.Printf("%s: %+v", unknwonError, err))
			return &ProductService.GetProductDetailsResponse{}, status.Errorf(
				codes.Unknown,
				unknwonError)
		}
	}

	return &ProductService.GetProductDetailsResponse{
		ProductName: result.Name,
		Desciption:  result.Description,
		Source:      result.Url,
		Type:        result.Type,
		Tags:        result.Tags,
		Creator:     result.Creator}, nil
}
