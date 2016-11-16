package controllers

import (
	"encoding/xml"
	//"fmt"
	"github.com/astaxie/beego"
	//"regexp"
	//"strings"
	"zhuiju365/wx"
)

//基础的消息类型
type WxMsg struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   string `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
}

//回复消息的子属性结构
//子属性图片信息
type Image struct {
	MediaId string `xml:"MediaId"`
}

//子属性语音信息
type Voice struct {
	MediaId string `xml:"MediaId"`
}

//子属性视频信息
type Video struct {
	MediaId     string `xml:"MediaId"`
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
}

//子属性音乐信息
type Music struct {
	Title        string `xml:"Title"`
	Description  string `xml:"Description"`
	HQMusicUrl   string `xml:"HQMusicUrl"`
	ThumbMediaId string `xml:"ThumbMediaId"`
}

//子属性图文回复项
type Item struct {
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	PicUrl      string `xml:"PicUrl"`
	Url         string `xml:"Url"`
}

//子属性图文回复集合
type Articles struct {
	Item []Item `xml:"item"`
}

//回复消息结构
//回复文本消息
type WxSeTextMsg struct {
	XMLName xml.Name `xml:"xml"`
	WxMsg
	Content string `xml:"Content"`
}

//回复图片消息
type WxSeImageMsg struct {
	XMLName xml.Name `xml:"xml"`
	WxMsg
	Image Image `xml:"Image"`
}

//回复语音消息
type WxSeVoiceMsg struct {
	XMLName xml.Name `xml:"xml"`
	WxMsg
	Voice Voice `xml:"Voice"`
}

//回复视频消息
type WxSeVideoMsg struct {
	XMLName xml.Name `xml:"xml"`
	WxMsg
	Video Video `xml:"Video"`
}

//回复音乐消息
type WxSeMusicMsg struct {
	XMLName xml.Name `xml:"xml"`
	WxMsg
	Music Music `xml:"Music"`
}

//回复图文消息
type WxSeArticlesMsg struct {
	XMLName xml.Name `xml:"xml"`
	WxMsg
	ArticleCount int32    `xml:"ArticleCount"`
	Articles     Articles `xml:"Articles"`
}

//收到信息，临时用
//收到的消息类型
type WxReceiveMsg struct {
	WxMsg
	MsgId string `xml:"MsgId"`
}

//接收的文本类型消息
type WxReTextMsg struct {
	XMLName xml.Name `xml:"xml"`
	WxReceiveMsg
	Content string `xml:"Content"`
}
type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	xml_info := "<xml><ToUserName><![CDATA[gh_33d7dbaf9e0c]]></ToUserName><FromUserName><![CDATA[o6nFtwGKAV4SvkzU50iIKwFa8gcc]]></FromUserName><CreateTime>1478854521</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[工]]></Content><MsgId>6351631803664816440</MsgId></xml>"
	/*
		//处理消息类型
		reg := regexp.MustCompile("<MsgType><!\\[CDATA\\[(.+)\\]\\]></MsgType>")
		xml_type := reg.FindString(xml_info)
		xml_type = xml_type[strings.Index(xml_type, "[")+1 : strings.LastIndex(xml_type, "]")]
		xml_type = xml_type[strings.Index(xml_type, "[")+1 : strings.LastIndex(xml_type, "]")]

		//把xml转换为struct
		reText := WxReTextMsg{}
		xml.Unmarshal([]byte(xml_info), &reText)
		this.Data["msg"] = xml_info
		fmt.Println(reText)
	*/

	recmsg := wx.InitRecMsg(xml_info)
	rep_str := wx.ReplyText("测试回复", recmsg)
	this.Data["sen"] = rep_str
	this.Data["rep"] = xml_info
	this.Data["Token"] = beego.AppConfig.String("wxToken")
	this.TplName = "index.html"
}
