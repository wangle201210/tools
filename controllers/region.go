package controllers

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"github.com/wangle201210/tools/models"
)

type RegionController struct {
	BaseController
}

type data struct {
	Addr string
	Birth string
	Gender string
	Err string
}
func (c *RegionController) URLMapping() {
	c.Mapping("Get", c.Get)
}

// @router /region [get]
func (this *RegionController) Get() {
	var info data
	idcard := this.GetString("idcard")
	if len(idcard) == 18 {
		// 取出前6位用来查找地址
		i, err := strconv.ParseInt(string([]byte(string(idcard)[0:6])), 10, 64)
		if err != nil {
			info.Err = "身份证号不合法！"
		}
		if info.Addr, err = getAddr(i); err != nil {
			info.Err = "身份证号不合法！"
		}
		info.Birth = getBirth(idcard)
		i2, _ := strconv.ParseInt(string([]byte(string(idcard)[16:17])), 10, 64)
		if i2&1 == 1 {
			info.Gender = "男"
		} else {
			info.Gender = "女"
		}
	} else if(len(idcard)) ==0 {
		info.Err = "请输入身份证号！"
	} else {
		info.Err = "身份证号不合法！"
	}
	data := Response{200,"success",info}
	this.Data["json"] = data
	this.ServeJSON()
	return
}

func getBirth(s string) (res string) {
	res = string([]byte(string(s)[6:10])) + "年" +
		string([]byte(string(s)[10:12])) + "月" +
		string([]byte(string(s)[12:14])) + "日"
	return
}

func getAddr(idcard int64) (string,error) {
	res := ""
	o := orm.NewOrm()
	self := models.Region{Id: idcard}
	read := o.Read(&self)
	if read != nil {
		return "",read
	}
	p := models.Region{Id: self.ParentId}
	read_parent := o.Read(&p)
	if read_parent != nil {
		return "",read
	}
	pp := models.Region{Id: p.ParentId}
	read_pp := o.Read(&pp)
	if read_pp != nil {
		return "",read
	}
	res = pp.Name + "省" + p.Name + self.Name
	return res,nil
}

func init() {
}