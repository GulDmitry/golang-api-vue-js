package controllers

import (
	"time"
	"log"
	"github.com/satori/go.uuid"
	"github.com/guldmitry/go-api-vue-js/controllers/rest"
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego"
	"github.com/guldmitry/go-api-vue-js/models"
)

type TaskController struct {
	MainController
}

type Task struct {
	Id    uuid.UUID `form:"-"`
	Title string    `form:"title"`
	Body  string    `form:"body"`
	Date  time.Time
}

func (c *TaskController) Index() {
	// Render Form https://beego.me/docs/mvc/view/view.md#renderform
	c.TplName = "new.tpl"
	// Autorender is disabled.
	c.Render();
}

func (c *TaskController) Post() {
	t := Task{}
	if err := c.ParseForm(&t); err != nil {
		beego.Info("Form parse error!")
		c.Abort("401")
	}

	valid := validation.Validation{}
	valid.Required(t.Title, "name")
	valid.MaxSize(t.Body, 70, "bodyMax")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		// TODO: show errors.
		c.Render();
	} else {
		task, _ := models.NewTask(t.Title, t.Body)
		rest.TaskManager.Save(task)

		c.Redirect("/", 302)
	}
}
