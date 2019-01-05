package main

import (
	"github.com/astaxie/beego"
	"tools/models"
	_ "tools/routers"
)

func main() {
	beego.Run()
}

func init() {
	models.Init()
	beego.SetStaticPath("/static","static")
}

