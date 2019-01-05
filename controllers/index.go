package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) URLMapping() {
	c.Mapping("Get", c.Get)
}

// @router / [get]
func (this *IndexController) Get() {
	this.TplName = "index.html"
}


func init() {
	beego.BConfig.WebConfig.TemplateLeft="<<<"
	beego.BConfig.WebConfig.TemplateRight=">>>"
}