package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	// Anonymous field, so the MainController has all methods that beego.Controller has.
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Title"] = "Golang + VueJs"
	// If not specified, tried to find maincontroller/get.tpl
	c.TplName = "index.tpl"
	// Autorender is disabled.
	c.Render();
}
