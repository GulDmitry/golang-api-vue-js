package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/guldmitry/go-api-vue/models"
)

// Operations about Tasks
type TaskController struct {
	beego.Controller
}

// @Title GetAll
// @Description Get all Tasks
// @Success 200 {object} models.Task
// @router / [get]
func (this *TaskController) GetAll() {
	res := struct{ Tasks []*models.Task }{models.DefaultTaskList.All()}
	this.Data["json"] = res
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
	beego.Info("Task is ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	t, ok := models.DefaultTaskList.Find(intid)
	beego.Info("Found", ok)
	if !ok {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("task not found"))
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
	req := struct{ Title string }{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("empty title"))
		return
	}
	t, err := models.NewTask(req.Title)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	models.DefaultTaskList.Save(t)
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
	beego.Info("Task is ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	var t models.Task
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &t); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	if t.Id != intid {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("inconsistent task Ids"))
		return
	}
	if _, ok := models.DefaultTaskList.Find(intid); !ok {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("task not found"))
		return
	}
	models.DefaultTaskList.Save(&t)
}

// @Title Delete
// @Description delete the task
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *TaskController) Delete() {
	id := this.GetString(":id")
	intid, _ := strconv.ParseInt(id, 10, 64)
	models.DefaultTaskList.Delete(intid)
	this.Data["json"] = "delete success!"
	this.ServeJSON()
}
