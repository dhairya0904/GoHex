package dao

import (
	"context"
	"rotteneggs/productservice/common/constants"
	"testing"
)

type mockFireStoreClient struct {
}

var (
	fireStoreMockCall func(id string) (map[string]interface{}, error)
)

func (mockFireStoreClient *mockFireStoreClient) Get(id string) (map[string]interface{}, error) {
	return fireStoreMockCall(id)
}

func TestGet(t *testing.T) {
	fireStoreMockCall = func(id string) (map[string]interface{}, error) {
		return constants.ProductDetailsMap, nil
	}
	fireStore := GetProductDataFireBase(&mockFireStoreClient{})
	actual, err := fireStore.Get("testId")
	if err != nil {
		t.Errorf(constants.UnexpectedError)
	}
	if actual.Name != constants.Title {
		t.Errorf("Got wrong name")
	}
	if actual.Type != constants.Type {
		t.Errorf("Got wrong type")
	}
	if actual.Url != constants.SourceUrl {
		t.Errorf("Got wrong type")
	}
}

func TestGetForTimeLimitExceeded(t *testing.T) {
	fireStoreMockCall = func(id string) (map[string]interface{}, error) {
		return nil, context.DeadlineExceeded
	}
	fireStore := GetProductDataFireBase(&mockFireStoreClient{})
	_, err := fireStore.Get("testId")

	if err == nil {
		t.Errorf(constants.ExceptionShouldBeThrown)
	}
}
