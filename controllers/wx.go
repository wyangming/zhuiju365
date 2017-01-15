package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"zhuiju365/models"
	"zhuiju365/util"
	"zhuiju365/wx"
)

const (
	Email = "479027247@qq.com"
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
	if recmsg.MsgType == wx.MsgTypeEvent {
		if recmsg.Event == wx.EventSubscribe && recmsg.EventKey == "" {
			send_msg := "不要回复会电视剧名字跟会员之类的信息，平台不存那些东西，要账号信息可以直接查看历史消息或者等到第二天发新的会员信息。小编是个人兼职做这个公众号的，时间跟经历都非常有限。如果有朋友有兴趣一块儿做的可以直接联系我，邮箱：479027247@qq.com；"
			rep_str = wx.ReplyText(send_msg, recmsg)
			this.Ctx.WriteString(rep_str)
			return
		}
	}
	//如果不是文本类型的处理
	if recmsg.MsgType != wx.MsgTypeText {
		rep_str = wx.ReplyText("认别不了此类信息！建议与联系邮箱："+Email, recmsg)
		this.Ctx.WriteString(rep_str)
		return
	}
	//如果不是自己的处理
	/*
		if recmsg.FromUserName != "o6nFtwGKAV4SvkzU50iIKwFa8gcc" {
			rep_str = wx.ReplyText("距离新版本上线还有2天，请耐心等待！建议与联系邮箱："+Email, recmsg)
			this.Ctx.WriteString(rep_str)
			return
		}
	*/
	//处理要找的资源信息
	videos, err := models.FindVideo(recmsg.Content)
	//如果没有找到资源
	if err != nil || len(videos) < 1 {
		beego.Error(err)
		rep_str = wx.ReplyText("亲，没有找到这个资源,您可以换个别的名子试试，如果还没有的话，等有的话会第一时间通知你!建议与联系邮箱："+Email, recmsg)
		this.Ctx.WriteString(rep_str)
		return
	}
	base_url_str := util.BaseUrl(this.Ctx)
	//设置为9是为了最后一个广告
	article_count := 9
	if len(videos) < 9 {
		article_count = len(videos)
	}
	//新闻集合
	articles := make([]wx.Article, 0)
	for i := 0; i < article_count; i++ {
		video := videos[i]
		wx_url := fmt.Sprintf("%s/video/film/%d", base_url_str, video.Id)
		article := wx.Article{
			Title:       video.FilmName,
			Description: "",
			PicUrl:      "",
			Url:         wx_url,
		}
		articles = append(articles, article)
	}
	//最后沟通内容
	article_count++
	article := wx.Article{
		Title:       "建议与联系邮箱：" + Email,
		Description: "",
		PicUrl:      "",
		Url:         "",
	}
	articles = append(articles, article)
	rep_str = wx.ReplyNews(articles, recmsg)
	this.Ctx.WriteString(rep_str)
}
