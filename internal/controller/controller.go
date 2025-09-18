package controller

import (
	"dytui/internal/awsutil"
	"dytui/internal/config"
	"dytui/internal/dynamo"
	"errors"
)

var (
	ErrEmptyCredentials = errors.New("credentials is empty")
)

type Controller struct {
	currentProfile string
	currentSession *dynamo.Session
}

func New(awsProfile string) (*Controller, error) {

	cred, err := awsutil.LoadProfile(awsProfile)

	if err != nil {
		return nil, err
	}

	awsConfig, err := config.New(cred.Name, cred.Region)

	if err != nil {
		return nil, err
	}

	ctrl := &Controller{
		currentProfile: awsProfile,
		currentSession: dynamo.New(awsConfig),
	}

	return ctrl, nil
}

func (c *Controller) Current() *dynamo.Session {
	return c.currentSession
}
