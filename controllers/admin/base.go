// baseController
package admin

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/logs"
)

//子控制器初始化公共接口
type ControllerPreparer interface {
	ControllerPrepare()
}



//后台控制器基类
type BaseController struct {
	beego.Controller
}

func (b *BaseController) Prepare() {

	//后台所有控制器初始化，可以进行auth认证，链接数据库。。。
	//写文件日志
	//logs.SetLogger(logs.AdapterFile,`{"filename":"logs/project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	//logs.EnableFuncCallDepth(true)
	//x:=5
	//y:=6
	//logs.Debug("测试日志信息:", x, " y:", y)
	//logs.Debug("my book is bought in the year of ", 2016)
	//logs.Info("this %s cat is %v years old", "yellow", 3)
	//logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
	//logs.Error(1024, "is a very", "good game")



	//子控制器初始化
	if app, ok := b.AppController.(ControllerPreparer); ok {
		app.ControllerPrepare()
	}

}
