package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	v := c.GetSession("test")
	if v == nil {
		c.SetSession("test", int(56))
		fmt.Printf("setsession %d\n", v)
	} else {
		fmt.Printf("getsession %d\n", v)
	}
}
