// @APIVersion 1.0.0
// @Title Test API
// @Description
package routers

import (
	"github.com/guldmitry/go-api-vue/controllers"
	"github.com/astaxie/beego"
)

// Go initializes packages and runs init() in every package,
func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/task",
			beego.NSInclude(
				&controllers.TaskController{},
			),
		),
	)
	beego.AddNamespace(ns)

	beego.Router("/", &controllers.MainController{})
}
