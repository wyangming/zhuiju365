package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"zhuiju365/models"
)

type VideoController struct {
	beego.Controller
}

func (this *VideoController) Film() {
	tid := this.Ctx.Input.Params()["0"]
	this.TplName = "video_flim.html"

	video, resources, err := models.FindVideoRe(tid)
	if err != nil || video == nil {
		beego.Error(err)
		return
	}
	this.Data["Video"] = video
	if resources == nil {
		return
	}
	//处理一下百度云盘的内容
	for _, resource := range resources {
		if len(resource.BaiduYun) < 1 {
			continue
		}
		resource.BaiduYun = strings.Replace(resource.BaiduYun, "#", " 密码：", -1)
	}
	this.Data["Resources"] = resources
}
