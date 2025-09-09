package dynamo

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	ErrRegionNotDefined = errors.New("region is not defined")
)

type Session struct {
	client *dynamodb.Client
}

func New(ctx context.Context, profile, region string) (*Session, error) {

	if len(region) == 0 {
		region = os.Getenv("AWS_REGION")

		if len(region) == 0 {
			return nil, ErrRegionNotDefined
		}
	}

	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(region),
		config.WithSharedConfigProfile(profile),
	)

	if err != nil {
		return nil, err
	}

	return &Session{client: dynamodb.NewFromConfig(cfg)}, nil
}

func (s *Session) Scan(ctx context.Context, tableName string) ([]map[string]any, error) {

	output, err := s.client.Scan(ctx, &dynamodb.ScanInput{
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

func (s *Session) Query(ctx context.Context, tableName string) ([]map[string]any, error) {

	output, err := s.client.Scan(ctx, &dynamodb.ScanInput{
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

func (s *Session) ListTables(ctx context.Context) ([]string, error) {

	result, err := s.client.ListTables(ctx, &dynamodb.ListTablesInput{})

	if err != nil {
		return nil, err
	}

	return result.TableNames, nil
}
