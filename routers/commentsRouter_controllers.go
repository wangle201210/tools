package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["tools/controllers:IndexController"] = append(beego.GlobalControllerRouter["tools/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tools/controllers:RegionController"] = append(beego.GlobalControllerRouter["tools/controllers:RegionController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/region`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tools/controllers:TaxController"] = append(beego.GlobalControllerRouter["tools/controllers:TaxController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/tax`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
