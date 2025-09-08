package dynamo

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	dynamodbClient *dynamodb.Client
)

func init() {

	resolver := aws.EndpointResolverWithOptionsFunc(
		func(service, region string, _ ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:           "http://localhost:4566", // LocalStack
				SigningRegion: "us-east-1",
			}, nil
		})

	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolverWithOptions(resolver),
	)

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	dynamodbClient = dynamodb.NewFromConfig(cfg)
}

func Scan(ctx context.Context, tableName string) ([]map[string]any, error) {

	output, err := dynamodbClient.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String("Employees"),
	})

	if err != nil {
		return nil, err
	}

	items := output.Items

	result := make([]map[string]any, 0)

	for _, item := range items {

		var m map[string]any

		if err := attributevalue.UnmarshalMap(item, &m); err != nil {
			return nil, err
		}

		result = append(result, m)
	}

	return result, nil
}

func ListTables(ctx context.Context) ([]string, error) {

	result, err := dynamodbClient.ListTables(ctx, &dynamodb.ListTablesInput{})

	if err != nil {
		return nil, err
	}

	return result.TableNames, nil
}
