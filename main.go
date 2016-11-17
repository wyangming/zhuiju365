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
	_ "zhuiju365/routers"
)

//2344#$
func main() {
	m1 := make(map[int]int)
	m1[1] = 3
	m1[8] = 10

	for k, v := range m1 {
		fmt.Println(fmt.Sprintf("this key is %d value is %d", k, v))
	}

	web()
}
func initData() {
	for i := 210; i <= 213; i++ {
		count := strconv.Itoa(i)
		open_file_path := "bdfilm" + string(filepath.Separator) + "bd_file" + count + ".txt"
		file, err := os.Open(open_file_path)
		if err != nil {
			beego.Error(err)
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
	//微信接口
	beego.Router("/wx", &controllers.WxController{})
	beego.AutoRouter(&controllers.VideoController{})
}
