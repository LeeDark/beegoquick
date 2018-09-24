// @APIVersion 1.0
// @Title Beego Quick Solutions
// @Description Here will be description about project
// @Contact server.povarnin@gmail.com
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
