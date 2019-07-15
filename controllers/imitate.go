package controllers

import (
	"github.com/wangle201210/tools/pkg/worker"
)

type ImitateController struct {
	BaseController
}

func (c *ImitateController) URLMapping() {
	c.Mapping("Get", c.Get)
}

type imitate struct {
	url string
}
// @router /imitate [get]
func (this *ImitateController) Get() {
	s := this.GetString("url")
	filePath := worker.Imitate(s)
	data := Response{200,"success",filePath}
	this.Data["json"] = data
	this.ServeJSON()
	return
}