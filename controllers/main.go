package controllers

import (
	"github.com/astaxie/beego"
	"github.com/guldmitry/go-api-vue-js/models"
	"encoding/json"
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MainController struct {
	// Anonymous field, so the MainController has all methods that beego.Controller has.
	beego.Controller
}

var demoData = []*models.Task{
	{"e977bc4d-ee93-4f98-a03f-d96734e042ba", "Demo title 1", "Body 2", time.Now()},
	{"b074ea11-6aec-4ea9-92c4-b2e473107244", "Demo title 2", "Body 1", time.Now()},
}

// DB Model.
type User struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}
type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`
}

func testDB() {
	// DB usage example
	orm.DefaultTimeLoc = time.UTC
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	// Model
	orm.RegisterModel(new(User), new(Post))

	// Set default database.
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(172.16.238.10:3306)/gavj?charset=utf8", 30)

	// Drop table and re-create.
	force := false
	// Print log.
	verbose := true
	// Error.
	err := orm.RunSyncdb("default", force, verbose)
	if err != nil {
		beego.Error("Cannot sync DB.")
	}

	o := orm.NewOrm()

	user := User{Name: "User 1"}
	post1 := Post{Title: "Post 1", User: &user}
	post2 := Post{Title: "Post 2", User: &user}

	// insert
	id, err := o.Insert(&user)
	log.Printf("ID: %d, ERR: %v\n", id, err)
	o.Insert(&post1)
	o.Insert(&post2)

	// update
	user.Name = "astaxie"
	num, err := o.Update(&user)
	log.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := User{Id: user.Id}
	err = o.Read(&u)
	log.Printf("USER: %v, ERR: %v\n", u, err)

	var posts []*Post
	num, errPost := o.QueryTable("post").Filter("User", user.Id).RelatedSel().All(&posts)
	if errPost == nil {
		log.Printf("%d posts read\n", num)
		for _, post := range posts {
			log.Printf("Id: %d, UserName: %s, Title: %s\n", post.Id, post.User.Name, post.Title)
		}
	}

	var userRes User
	errUser := o.QueryTable("user").Filter("posts__title", "Post 1").Limit(1).One(&userRes)
	if errUser == nil {
		log.Printf("User %v", userRes)
	}

	o.LoadRelated(&user, "Posts")
	log.Printf("User with posts loaded: %v", user.Posts)

	// delete
	num, err = o.Delete(&u)
	log.Printf("NUM: %d, ERR: %v\n", num, err)
}

func init() {
	beego.Debug("Main Controller init")
	// Uncomment to test ORM.
	//testDB()
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
