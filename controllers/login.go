package controllers

import (
	"github.com/astaxie/beego"
	"strings"
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
	this.TplName = "login.html"
	username := this.Input().Get("username")
	username = strings.Trim(username, " ")
	if len(username) < 1 {
		this.Data["message"] = "用户名不可以为空"
		return
	}
	password := this.Input().Get("password")
	password = strings.Trim(password, " ")
	if len(password) < 1 {
		this.Data["message"] = "密码不可以为空"
		return
	}
	conf_loginuser := beego.AppConfig.String("loginuser")
	conf_loginpwd := beego.AppConfig.String("loginpwd")

	if username != conf_loginuser || password != conf_loginpwd {
		this.Data["message"] = "用户名密码不对"
		return
	}
}
