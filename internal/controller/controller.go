package controller

import "dytui/internal/dynamo"

type Controller struct {
	sessions       map[string]dynamo.Session
	currentSession dynamo.Session
}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) List() {

}

func (c *Controller) Switch() {

}

func (c *Controller) Current() {

}

func getConnection() {

}
