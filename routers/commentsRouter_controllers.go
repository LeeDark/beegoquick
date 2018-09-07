package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/LeeDark/beegoquick/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/LeeDark/beegoquick/controllers:UserController"],
		beego.ControllerComments{
			Method: "Hello",
			Router: `/users`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/LeeDark/beegoquick/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/LeeDark/beegoquick/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetUser",
			Router: `/users/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("id", param.InPath),
				param.New("name"),
			),
			Params: nil})

}
