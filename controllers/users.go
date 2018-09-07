package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) URLMapping() {
	this.Mapping("Hello", this.Hello)
}

// @router /users
func (this *UserController) Hello() {
	this.Data["Hello"] = "Hello, Sergey!"
	this.Data["ID"] = 505
	this.TplName = "user.tpl"
	this.Render()
}

// /users/:id?name=...
// @router /users/:id
func (this *UserController) GetUser(id int, name string) {
	this.Data["Hello"] = "Hello, " + name + "!"
	this.Data["ID"] = id
	this.TplName = "user.tpl"
	this.Render()
}
