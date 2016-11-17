package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
	"time"
)

//视频信息
type Video struct {
	Id          int64
	FilmName    string    //视频名
	Title       string    //视频标题
	Attachment  string    //视频标签
	Type        string    //TV电视剧 FILM电影
	SysTime     time.Time //系统抓取的时间
	ReleaseDate time.Time `orm:"type(date)"` //上映时间
	UpTime      time.Time `orm:"type(date)"` //更新时间
	Fbak        string    //备注
	Vstatus     int       `orm:"default(1)"` //状态，0永久1可以需审核 可以删除
}

//电影电视剧资源信息
type VideoResource struct {
	Id          int64
	FilmId      int64  //电影电视剧的Id
	ReName      string `orm:"index"` //资源名称
	DownloadUrl string //磁力链接
	BaiduYun    string //百度云盘
	Vrbak       string //资源备注
}

//初始化数据库配置
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	mysqlurl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", beego.AppConfig.String("dbuser"), beego.AppConfig.String("dbpass"), beego.AppConfig.String("dburl"), beego.AppConfig.String("dbport"), beego.AppConfig.String("dbname"))
	orm.RegisterDataBase("default", "mysql", mysqlurl, 30, 60)
	orm.RegisterModel(new(Video), new(VideoResource))

	// 开启 ORM 调试模式
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)
}

//查询一个视频信息，带资源
func FindVideoRe(videoStr string) (*Video, []*VideoResource, error) {
	videoId, err := strconv.Atoi(videoStr)
	if err != nil {
		return nil, nil, err
	}
	o := orm.NewOrm()
	qs := o.QueryTable("video")
	video := &Video{}
	err = qs.Filter("id", videoId).One(video)
	//如果报错与没有找到
	if err != nil || len(video.FilmName) < 1 {
		return nil, nil, err
	}
	re_qs := o.QueryTable("video_resource")
	resources := make([]*VideoResource, 0)
	_, err = re_qs.Filter("film_id", videoId).All(&resources)
	if err != nil || len(resources) < 1 {
		return video, nil, err
	}
	return video, resources, nil
}

//查询一个视频信息，不带资源
func FindVideo(video_name string) ([]*Video, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("video")
	qs = qs.Filter("attachment__contains", "$"+video_name+"#")
	videos := make([]*Video, 0)
	_, err := qs.All(&videos)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func MysqlTest(str_films []string) {
	o := orm.NewOrm()
	nowTime := time.Now()
	for _, str_film := range str_films {
		films := strings.Split(str_film, "&")
		count := len(films)
		if count < 3 {
			continue
		}
		rutime, _ := time.Parse("2006-01-02", films[0])
		title := films[1]
		name := title[strings.Index(title, "《")+3 : strings.Index(title, "》")]
		video := &Video{
			FilmName:    name,
			Title:       title,
			Attachment:  "$" + name + "#",
			Type:        "FILM",
			SysTime:     nowTime,
			ReleaseDate: rutime,
			UpTime:      rutime,
		}
		vre := &VideoResource{
			ReName:      name,
			DownloadUrl: films[2],
		}
		if count > 5 {
			vre.BaiduYun = films[3] + "#" + films[4]
		}
		_, err := o.Insert(video)
		if err != nil {
			beego.Error(err)
			return
		}
		vre.FilmId = video.Id
		_, err = o.Insert(vre)
		if err != nil {
			beego.Error(err)
			return
		}
		fmt.Println(video)
		fmt.Println(vre)
	}

	/*
		row := &VideoResource{
			FilmId:   2,
			ReName:   "doc",
			Ed2k:     "ed2k",
			BaiduYun: "baiduyun",
			Vrbak:    "vrbak",
		}
		o := orm.NewOrm()
		o.Insert(row)
	*/
}
