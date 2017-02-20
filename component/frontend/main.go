package main

import (
	_ "Clio/component/frontend/routers"

	_ "Clio/component/backend/conf"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
