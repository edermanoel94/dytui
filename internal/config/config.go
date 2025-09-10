package config

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var (
	ErrRegionNotDefined  = errors.New("region is not defined")
	ErrProfileNotDefined = errors.New("profile is not defined")
)

type Config struct {
	awsConfig aws.Config
}

func New(profile, region string) (*Config, error) {

	if len(region) == 0 {
		region = os.Getenv("AWS_REGION")

		if len(region) == 0 {
			return nil, ErrRegionNotDefined
		}
	}

	if len(profile) == 0 {
		return nil, ErrProfileNotDefined
	}

	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(region),
		config.WithSharedConfigProfile(profile),
	)

	if err != nil {
		return nil, err
	}

	return &Config{awsConfig: cfg}, nil
}

func (c *Config) AWS() aws.Config {
	return c.awsConfig
}
