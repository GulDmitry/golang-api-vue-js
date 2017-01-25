package controllers

import (
	"github.com/astaxie/beego"
	"github.com/guldmitry/go-api-vue-js/models"
	"encoding/json"
	"time"
)

type MainController struct {
	// Anonymous field, so the MainController has all methods that beego.Controller has.
	beego.Controller
}

var demoData = []*models.Task{
	{"e977bc4d-ee93-4f98-a03f-d96734e042ba", "Demo title 1", "Body 2", time.Now()},
	{"b074ea11-6aec-4ea9-92c4-b2e473107244", "Demo title 2", "Body 1", time.Now()},
}

func (c *MainController) Prepare() {
	beego.Debug("Main controller Prepare.")
	if beego.AppConfig.String("runmode") == "dev" {
		c.Data["assetsUrl"] = "http://localhost:8081/"
	} else {
		c.Data["assetsUrl"] = ""
	}
	// To receive flash messages.
	beego.ReadFromRequest(&c.Controller)

	// Session specific set of tasks.
	manager := models.NewTaskManager()
	stm := c.GetSession("TaskManager")
	if stm == nil {
		// Populate demo data.
		for _, v := range demoData {
			// Do not use []models.Task and &v. It uses one pointer
			manager[v.Id] = v
		}
		// Saving models.Tasks in session directly causes error on server restart:
		// gob: name not registered for interface: "github.com/guldmitry/go-api-vue-js/models.Tasks"
		v, err := json.Marshal(manager)
		if err != nil {
			beego.Error(err)
		}
		c.SetSession("TaskManager", v)
	} else {
		json.Unmarshal(stm.([]byte), &manager)
	}
	TaskManager = manager

	c.Data["Title"] = "Golang + VueJs"
	c.Layout = "layout.tpl"

	c.LayoutSections = map[string]string{
		"ErrorBox":"error_box.tpl",
	}
	// If not specified, tried to find maincontroller/get.tpl
	c.TplName = "index.tpl"
}

// Dump tasks to session.
func (c *MainController) Finish() {
	beego.Info("Dump TaskManager to session")
	v, _ := json.Marshal(TaskManager)
	c.SetSession("TaskManager", v)
}

func (c *MainController) Get() {
	// Autorender is disabled.
	c.Render();
}
