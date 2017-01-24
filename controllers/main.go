package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	// Anonymous field, so the MainController has all methods that beego.Controller has.
	beego.Controller
}
func (c *MainController) Prepare() {
	if beego.AppConfig.String("runmode") == "dev" {
		c.Data["assetsUrl"] = "http://localhost:8081/"
	} else {
		c.Data["assetsUrl"] = ""
	}
	// To receive flash messages.
	beego.ReadFromRequest(&c.Controller)

	c.Data["Title"] = "Golang + VueJs"
	c.Layout = "layout.tpl"

	c.LayoutSections = map[string]string{
		"ErrorBox":"error_box.tpl",
	}
	// If not specified, tried to find maincontroller/get.tpl
	c.TplName = "index.tpl"
}

func (c *MainController) Get() {
	// Autorender is disabled.
	c.Render();
}
