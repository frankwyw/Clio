package routers

import (
	"Clio/component/frontend/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/time", &controllers.TimeController{}, "get:Query")

	beego.Router("/registry/login", &controllers.RegController{}, "get:Reglogin")
	beego.Router("/registry/listrepo", &controllers.RegController{}, "get:RegListRepo")
	beego.Router("/registry/getimagetag", &controllers.RegController{}, "get:RegGetImageTag")
	beego.Router("/registry/push", &controllers.RegController{}, "get:RegPush")
	beego.Router("/registry/delimage", &controllers.RegController{}, "get:RegDelImage")
	beego.Router("/registry/pull", &controllers.RegController{}, "get:RegPull")

	beego.Router("/kubernetes", &controllers.K8sController{})

	beego.Router("/login", &controllers.LoginController{}, "get:Login")

	beego.Router("/prometheus", &controllers.PromController{}, "get:PromQuery")

}
