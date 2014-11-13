package routers

import (
	//	"conf"
	"controllers"
	"github.com/astaxie/beego"
	"model"
	//	"github.com/ozonesurfer/pastemgomodel/src/model"
)

func init() {
	model.Init("mgopastebintemp-2")
	beego.SetStaticPath("static", "C:/mgopastebin_temp/src/static")
	beego.Router("/", &controllers.MainController{})
	beego.Router("/new", &controllers.CreateController{})
	beego.Router("/paste/:id", &controllers.ShowController{})
}
