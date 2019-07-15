package main

import (
	"github.com/astaxie/beego"
	"github.com/wangle201210/tools/models"
	_ "github.com/wangle201210/tools/routers"
)

func main() {
	beego.Run()
}

func init() {
	models.Init()
	beego.SetStaticPath("/static","static")
}

