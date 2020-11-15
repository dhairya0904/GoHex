package dao

import (
	"context"
	config "rotteneggs/productservice/pkg/config"
	constants "rotteneggs/productservice/pkg/constants"
	logger "rotteneggs/productservice/pkg/logging"

	"cloud.google.com/go/firestore"
)

type FirestoreClientInterface interface {
	Get(id string) (map[string]interface{}, error)
}

type FirestoreClient struct{}

const (
	FireStoreClientErrorString string = "Error while creating firestore client"
	FireStoreQueryErrorString  string = "Error while while querying firestore"
)

func (firestoreClient *FirestoreClient) Get(id string) (map[string]interface{}, error) {
	log := logger.NewLogger()
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, config.DataBaseTimeOut)
	projectID := constants.ProjectID
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Error(FireStoreClientErrorString, err)
		return nil, err
	}

	defer client.Close()

	data, err := client.Collection(constants.ProductCollection).Doc(id).Get(ctx)
	if err != nil {
		log.Error(FireStoreQueryErrorString, err)
		return nil, err
	}
	return data.Data(), nil
}
