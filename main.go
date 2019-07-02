package main

import (
	_ "LightOfData/routers"

	"github.com/astaxie/beego"
)

func main() {
	//开启session
	beego.BConfig.WebConfig.Session.SessionOn = true
	//客户端  cookie
	beego.BConfig.WebConfig.Session.SessionName = "light_of_data_sessionid"
	//启动http监听
	beego.Run()
}
