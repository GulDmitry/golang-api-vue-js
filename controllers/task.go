package controllers

import (
	"time"
	"log"
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego"
	"github.com/guldmitry/go-api-vue-js/models"
)

type TaskController struct {
	MainController
}

type Task struct {
	Id    string    `form:"-"`
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

// @router /tasks/edit/:id [GET]
// @router /tasks/edit/:id [POST]
func (c *TaskController) Edit() {
	uid := c.Ctx.Input.Param(":id")
	beego.Info("Edit Task ", uid)

	t, ok := TaskManager.Find(uid)
	if !ok {
		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte("Task not found."))
		return
	}

	// Handle POST
	if (len(c.Input()) > 0) {
		flash := beego.NewFlash()
		taskForm := Task{}
		if err := c.ParseForm(&taskForm); err != nil {
			beego.Info("Form parse error!")
		}
		newTask, _ := models.NewTask(taskForm.Title, taskForm.Body)
		updatedTask, err := TaskManager.Update(uid, newTask)
		if err != nil {
			flash.Error("Edit failed.<br/>" + err.Error())
		} else {
			flash.Notice("Edited successfully.")
		}
		flash.Store(&c.Controller)

		c.Data["task"] = updatedTask
	} else {
		c.Data["task"] = t
	}

	c.TplName = "edit.tpl"
	c.Render();
}

func (c *TaskController) Post() {
	flash := beego.NewFlash()
	t := Task{}
	if err := c.ParseForm(&t); err != nil {
		beego.Info("Form parse error!")
		c.Abort("401")
	}

	valid := validation.Validation{}
	valid.Required(t.Title, "name")
	valid.MaxSize(t.Body, 70, "bodyMax")

	if valid.HasErrors() {
		errorMessage := ""
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			errorMessage += err.Message + "<br/>"

		}
		flash.Error(errorMessage)
		flash.Store(&c.Controller)
		c.Index();
		return
	} else {
		task, _ := models.NewTask(t.Title, t.Body)
		TaskManager.Save(task)

		flash.Notice("Task was created successfully.")
		flash.Store(&c.Controller)

		c.Redirect("/", 302)
	}
}
