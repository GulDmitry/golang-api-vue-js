package routers

import (
	"github.com/guldmitry/go-api-vue/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
