package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greet := "Fuck you world"
	return c.Render(greet)
}
