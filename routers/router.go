package routers

import (
	"github.com/LeeDark/beegoquick/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.UserController{})
}
