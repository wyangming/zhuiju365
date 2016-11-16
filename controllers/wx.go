package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"zhuiju365/models"
	"zhuiju365/wx"
)

type WxController struct {
	beego.Controller
}

//微信认证
func (this *WxController) Get() {
	timestamp, nonce, signatureIn, echostr := this.GetString("timestamp"), this.GetString("nonce"), this.GetString("signature"), this.GetString("echostr")
	this.Ctx.WriteString(wx.WxAuth(timestamp, nonce, signatureIn, echostr))
}

//消息接收
func (this *WxController) Post() {
	var str_msg string = string(this.Ctx.Input.RequestBody)
	var rep_str string
	recmsg := wx.InitRecMsg(str_msg)
	//如果不是文本类型的处理
	if recmsg.MsgType != wx.MsgTypeText {
		rep_str = wx.ReplyText("认别不了此类信息！", recmsg)
		this.Ctx.WriteString(rep_str)
	}
	//如果不是自己的处理
	if recmsg.FromUserName != "o6nFtwGKAV4SvkzU50iIKwFa8gcc" {
		rep_str = wx.ReplyText("距离新版本上线还有2天，请耐心等待！", recmsg)
		this.Ctx.WriteString(rep_str)
		return
	}
	//处理要找的资源信息
	video, resource, err := models.FindVideo(recmsg.Content)
	//如果没有找到资源
	if err != nil {
		rep_str = wx.ReplyText("亲，没有找到这个资源，等有的话会第一时间通知你", recmsg)
		this.Ctx.WriteString(rep_str)
	}
	resource_str := resource.BaiduYun
	resource_str = strings.Replace(resource_str, "#", "密码：", -1)
	rep_str = wx.ReplyText(video.FilmName+"  "+resource_str, recmsg)

	this.Ctx.WriteString(rep_str)
}