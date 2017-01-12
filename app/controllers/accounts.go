package controllers

import "github.com/revel/revel"

type Accounts struct {
	*revel.Controller
}

func (c Accounts) Create() revel.Result {
	return c.Render()
}

func (c Accounts) Payment() revel.Result {
	return c.Render()
}
