package aws

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBClient struct {
	client *dynamodb.Client
}

func NewDynamoDBClient() *DynamoDBClient {

	region := os.Getenv("AWS_REGION")

	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(region),
		//config.WithEndpointResolverWithOptions(),
	)

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return &DynamoDBClient{
		client: dynamodb.NewFromConfig(cfg),
	}
}

func (d DynamoDBClient) Client() *dynamodb.Client {
	return d.client
}
