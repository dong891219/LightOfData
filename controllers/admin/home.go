// home
package admin

import "html/template"

type HomeController struct {
	BaseController
}

func (h *HomeController) ControllerPrepare() {


}

func (h *HomeController) Index() {

	h.Data["Content"] = "hello world"

	h.Layout = "admin/layout.html"
	h.TplName = "admin/home/index.html"
}

func (h *HomeController)Login()  {
	h.Data["xsrfdata"]=template.HTML(h.XSRFFormHTML())
	h.TplName="admin/home/login.html"
}


