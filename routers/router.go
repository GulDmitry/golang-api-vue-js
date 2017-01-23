// @APIVersion 1.0.0
// @Title Test API
// @Description
package routers

import (
	"github.com/guldmitry/go-api-vue-js/controllers"
	restControllers "github.com/guldmitry/go-api-vue-js/controllers/rest"
	"github.com/astaxie/beego"
)

// Go initializes packages and runs init() in every package,
func init() {
	beego.SetStaticPath("/assets", "static/assets")

	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/tasks",
			beego.NSInclude(
				&restControllers.TaskController{},
			),
		),
	)
	beego.AddNamespace(ns)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/tasks/new", &controllers.TaskController{}, "GET:Index")
	beego.Router("/tasks/new", &controllers.TaskController{})
}
