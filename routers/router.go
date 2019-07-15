package routers

import (
	"github.com/wangle201210/tools/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.SetStaticPath("/files","files")

	beego.Include(&controllers.IndexController{})
	beego.Include(&controllers.RegionController{})
	beego.Include(&controllers.TaxController{})
	beego.Include(&controllers.ImitateController{})
}
