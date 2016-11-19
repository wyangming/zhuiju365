package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

//跳转到登录页面
func (this *LoginController) Get() {
	this.TplName = "login.html"
}

//登录方法
func (this *LoginController) Post() {

}
