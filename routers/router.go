package routers

import (
	"github.com/LeeDark/beegoquick/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/users",
				beego.NSInclude(
					&controllers.UserController{},
				),
			),
		)

	beego.AddNamespace(ns)
}
