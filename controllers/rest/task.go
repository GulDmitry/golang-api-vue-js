package rest

import (
	"github.com/astaxie/beego"
	"github.com/guldmitry/go-api-vue-js/models"
	"github.com/satori/go.uuid"
	"time"
	"encoding/json"
)

var TaskManager models.Tasks

func init() {
	TaskManager = models.NewTaskManager()
	t1 := models.Task{uuid.FromStringOrNil("e977bc4d-ee93-4f98-a03f-d96734e042ba"), "Title 1", "Body 2", time.Now()}
	t2 := models.Task{uuid.FromStringOrNil("b074ea11-6aec-4ea9-92c4-b2e473107244"), "Title 2", "Body 1", time.Now()}
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
func (c *TaskController) GetAll() {
	c.Data["json"] = TaskManager.All()
	c.ServeJSON()
}

// @Title Get
// @Description Get Task by Id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Task
// @Failure 403 :uid is empty
// @router /:id [get]
func (c *TaskController) Get() {
	id := c.Ctx.Input.Param(":id")
	uid := uuid.FromStringOrNil(id)
	beego.Info("Task is ", id, uid)

	t, ok := TaskManager.Find(uid)
	beego.Info("Found", ok)
	if !ok {
		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte("Task not found."))
		return
	}
	c.Data["json"] = t
	c.ServeJSON()
}

// @Title CreateTask
// @Description create tasks
// @Param	body		body 	models.Task	true		"body for task content"
// @Success 200 {int} models.Task.Id
// @Failure 403 body is empty
// @router / [post]
func (c *TaskController) Post() {
	req := struct {
		Title string
		Body  string
	}{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Empty title or body."))
		return
	}
	t, err := models.NewTask(req.Title, req.Body)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
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
func (c *TaskController) Put() {
	id := c.Ctx.Input.Param(":id")
	uid := uuid.FromStringOrNil(id)

	var t models.Task
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &t); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	beego.Info("Task is ", t)

	if _, ok := TaskManager.Find(uid); !ok {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Task not found."))
		return
	}

	// Generate a valid Id to pass validation check.
	t.Id = uuid.NewV4()

	_, err := TaskManager.Update(uid, t)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}
}

// @Title Delete
// @Description delete the task
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TaskController) Delete() {
	id := c.GetString(":id")
	uid := uuid.FromStringOrNil(id)

	TaskManager.Delete(uid)

	c.Data["json"] = "delete success!"
	c.ServeJSON()
}
