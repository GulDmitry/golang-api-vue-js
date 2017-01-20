package controllers

import (
	"github.com/astaxie/beego"
	"github.com/guldmitry/go-api-vue/models"
	"github.com/satori/go.uuid"
	"time"
	"encoding/json"
)

var TaskManager models.Tasks

func init() {
	TaskManager = models.NewTaskManager()
	t1 := models.Task{uuid.FromStringOrNil("e977bc4d-ee93-4f98-a03f-d96734e042ba"), "Title 1", "Body", time.Now()}
	t2 := models.Task{uuid.FromStringOrNil("b074ea11-6aec-4ea9-92c4-b2e473107244"), "Title 2", "Body", time.Now()}
	TaskManager[t1.Id] = &t1;
	TaskManager[t2.Id] = &t2;
}

// Operations about Tasks
type TaskController struct {
	beego.Controller
}

// @Title GetAll
// @Description Get all Tasks
// @Success 200 {object} models.Task
// @router / [get]
func (this *TaskController) GetAll() {
	this.Data["json"] = TaskManager.All()
	this.ServeJSON()
}

// @Title Get
// @Description Get Task by Id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Task
// @Failure 403 :uid is empty
// @router /:id [get]
func (this *TaskController) Get() {
	id := this.Ctx.Input.Param(":id")
	uid := uuid.FromStringOrNil(id)
	beego.Info("Task is ", id, uid)

	t, ok := TaskManager.Find(uid)
	beego.Info("Found", ok)
	if !ok {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("Task not found."))
		return
	}
	this.Data["json"] = t
	this.ServeJSON()
}

// @Title CreateTask
// @Description create tasks
// @Param	body		body 	models.Task	true		"body for task content"
// @Success 200 {int} models.Task.Id
// @Failure 403 body is empty
// @router / [post]
func (this *TaskController) Post() {
	req := struct {
		Title string
		Body  string
	}{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Empty title or body."))
		return
	}
	t, err := models.NewTask(req.Title, req.Body)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	TaskManager.Save(t)
}

// @Title Update
// @Description update the task
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Task	true		"body for task content"
// @Success 200 {object} models.Task
// @Failure 403 :id is not int
// @router /:id [put]
func (this *TaskController) Put() {
	id := this.Ctx.Input.Param(":id")
	uid := uuid.FromStringOrNil(id)

	var t models.Task
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &t); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	beego.Info("Task is ", t)

	if _, ok := TaskManager.Find(uid); !ok {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Task not found."))
		return
	}

	// Generate a valid Id to pass validation check.
	t.Id = uuid.NewV4()

	_, err := TaskManager.Update(uid, t)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}
}

// @Title Delete
// @Description delete the task
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *TaskController) Delete() {
	id := this.GetString(":id")
	uid := uuid.FromStringOrNil(id)

	TaskManager.Delete(uid)

	this.Data["json"] = "delete success!"
	this.ServeJSON()
}
