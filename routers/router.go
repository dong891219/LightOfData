package routers

import (
	"LightOfData/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {

	adminNs :=
		beego.NewNamespace("/admin",
			beego.NSRouter("/", &admin.HomeController{}, "get:Index"),
			beego.NSRouter("/index", &admin.HomeController{}, "get:Index"),
			beego.NSRouter("/login", &admin.HomeController{}, "get:Login"),

	)
	frontNs :=
		beego.NewNamespace("/wx",
			beego.NSRouter("/", &admin.HomeController{}, "get:Index"),
			beego.NSRouter("/index", &admin.HomeController{}, "get:Index"),

		)
	beego.AddNamespace(adminNs)
	beego.AddNamespace(frontNs)
}
