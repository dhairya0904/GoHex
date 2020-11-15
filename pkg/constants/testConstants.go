package constants

import (
	"rotteneggs/productservice/internal/core/domain"
)

const Creator string = "Raju rastogi"
const Description string = "hellp"
const Title string = "hello world"
const Type string = "Blog"
const SourceUrl string = "rastogi.com"

var ProductDetailsMap = map[string]interface{}{
	"creator":     Creator,
	"descirption": Description,
	"tags":        []string{"hello", "world"},
	"title":       Title,
	"type":        Type,
	"sourceUrl":   SourceUrl,
}

var ProductDetailsStruct = domain.Product{
	Creator:     Creator,
	Description: Description,
	Name:        Title,
	Type:        Type,
	Url:         SourceUrl,
	Tags:        []string{"hello", "world"},
	ID:          "testId",
}

const UnexpectedError string = "Unexpected error"
const ExceptionShouldBeThrown string = "Exception should have been thrown"
