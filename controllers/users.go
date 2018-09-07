package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) URLMapping() {
	this.Mapping("HelloSimple", this.HelloSimple)
}

// @router /users
func (this *UserController) HelloSimple() {
	this.Data["Hello"] = "Hello, Sergey!"
	this.Data["ID"] = 505
	this.TplName = "user.tpl"
}

// @router /users/:id
func (this *UserController) HelloCool(id int, name string) {
	this.Data["Hello"] = "Hello, " + name + "!"
	this.Data["ID"] = id
	this.TplName = "user.tpl"
}
