package routers

import (
	"tools/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.RegionController{})
}
