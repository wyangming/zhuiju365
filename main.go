package main

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"zhuiju365/controllers"
	"zhuiju365/models"
)

//2344#$
func main() {
	//initData()
	web()
}
func initData() {
	for i := 1; i <= 17; i++ {
		count := strconv.Itoa(i)
		open_file_path := "bdfilm" + string(filepath.Separator) + "bd_file" + count + ".txt"
		file, err := os.Open(open_file_path)
		if err != nil {
			beego.Error(err)
			return
		}
		defer file.Close()
		bfrd := bufio.NewReader(file)
		str_films := make([]string, 0)
		for {
			line, err := bfrd.ReadString('\n')
			if err == io.EOF {
				break
			}
			str_films = append(str_films, line)
		}
		fmt.Println("\nadd " + count + " file info over...\n")
		models.MysqlTest(str_films)
	}
}
func web() {
	retgRouter()
	beego.Run()
}
func retgRouter() {
	//默认路由
	beego.Router("/", &controllers.MainController{})
	//微信接口
	beego.Router("/wx", &controllers.WxController{})
	//视频资源页面
	beego.AutoRouter(&controllers.VideoController{})
	//登录页面
	beego.Router("/login", &controllers.LoginController{})
}
