package dynamo

import (
	"context"
	"dytui/internal/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Session struct {
	config *config.Config
}

func New(config *config.Config) *Session {
	return &Session{
		config: config,
	}
}

func (s *Session) Scan(ctx context.Context, tableName string) ([]map[string]any, error) {

	client := dynamodb.NewFromConfig(s.config.AWS())

	output, err := client.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(tableName),
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

func (s *Session) Query(ctx context.Context, tableName string) ([]map[string]any, error) {

	client := dynamodb.NewFromConfig(s.config.AWS())

	output, err := client.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(tableName),
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

func (s *Session) ListTables(ctx context.Context) ([]string, error) {

	client := dynamodb.NewFromConfig(s.config.AWS())

	result, err := client.ListTables(ctx, &dynamodb.ListTablesInput{})

	if err != nil {
		return nil, err
	}

	return result.TableNames, nil
}
