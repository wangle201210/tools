package models

import (
	"github.com/astaxie/beego/orm"
)

type Region struct {
	Id       int64     `json:"id" orm:"column(id);pk;"`
	Name    string    `json:"name" orm:"column(name);size(100)"`
	ParentId int64    `json:"parent_id" orm:"column(parent_id);size(10);"`
}

func init() {
	orm.RegisterModel(new(Region))
}
func Regions() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Region))
}

//创建城市
func CreateRegion(region Region) Region {
	o := orm.NewOrm()
	o.Insert(&region)
	return region
}

// 通过子集查找父级
func (reg Region) Parent(parent_id int64) error {
	if err := orm.NewOrm().Read(Region{Id: parent_id}); err != nil {
		return err
	}
	return nil
}

// Region database CRUD methods include Insert, Read, Update and Delete
func (reg *Region) Insert() error {
	if _, err := orm.NewOrm().Insert(reg); err != nil {
		return err
	}
	return nil
}

func (reg *Region) Read(fields ...string) error {
	if err := orm.NewOrm().Read(reg, fields...); err != nil {
		return err
	}
	return nil
}

func (reg *Region) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(reg, fields...); err != nil {
		return err
	}
	return nil
}

func (reg *Region) Delete() error {
	if _, err := orm.NewOrm().Delete(reg); err != nil {
		return err
	}
	return nil
}
