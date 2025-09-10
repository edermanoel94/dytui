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
	sessions       map[string]*dynamo.Session
	currentSession *dynamo.Session
}

func New() (*Controller, error) {

	credentials, err := awsutil.LoadAWSCredentials()

	if err != nil {
		return nil, err
	}

	if len(credentials) == 0 {
		return nil, ErrEmptyCredentials
	}

	ctrl := &Controller{
		sessions: make(map[string]*dynamo.Session),
	}

	for _, cred := range credentials {

		awsConfig, err := config.New(cred.Name, cred.Region)

		if err != nil {
			return nil, err
		}

		dynamoSession := dynamo.New(awsConfig)

		ctrl.sessions[cred.Name] = dynamoSession
	}

	return ctrl, nil
}

func (c *Controller) Switch() {

}

func (c *Controller) Current() *dynamo.Session {
	return c.currentSession
}
