package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

type AUser struct {
	Name string
	Age  int
}

func (this *UserController) URLMapping() {
	this.Mapping("Hello", this.Hello)
}

// @router /users
func (this *UserController) Hello() {
	this.Data["Hello"] = "Hello, Sergey!"
	this.Data["ID"] = 505
	this.Data["user"] = &AUser{Name: "Sergey", Age: 39}
	// default: usercontroller/hello.tpl
	this.TplName = "user.tpl"
	this.Render()
}

// /users/:id?name=...
// @router /users/:id
func (this *UserController) GetUser(id int, name string) {
	this.Data["Hello"] = "Hello, " + name + "!"
	this.Data["ID"] = id
	//this.Data["user"] = &AUser{Name: name, Age: id}
	this.Data["user"] = map[string]interface{}{"Name": name, "Age": id}
	this.TplName = "user.tpl"
	this.Render()
}
