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
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写用户名"}
		this.ServeJSON()
	}
	password := this.Input().Get("password")
	password = strings.Trim(password, " ")
	if len(password) < 1 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写密码"}
		this.ServeJSON()
	}
	conf_loginuser := beego.AppConfig.String("loginuser")
	conf_loginpwd := beego.AppConfig.String("loginpwd")

	if username != conf_loginuser || password != conf_loginpwd {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
		this.ServeJSON()
	}
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "贺喜你，登录成功"}
	this.ServeJSON()
}
