package controllers

type TaxController struct {
	BaseController
}

type info struct {
	top float64
	rates float64
	deduction float64
}

func (c *TaxController) URLMapping() {
	c.Mapping("Get", c.Get)
}


// @router /tax [get]
func (this *TaxController) Get() {
	salary, _ := this.GetFloat("salary")
	social, _ := this.GetFloat("social")
	oldStart := float64(3500)
	newStart := float64(5000)
	info := tax(salary, social, oldStart, newStart)
	data := Response{200,"success",info}
	this.Data["json"] = data
	this.ServeJSON()
	return
}


func tax(salary,social,oldStart,newStart float64) map[string]float64 {
	newLevels := []info{
		{3000,0.03,0},
		{12000,0.1,210},
		{25000,0.2,1410},
		{35000,0.25,2660},
		{55000,0.3,4410},
		{80000,0.35,7160},
		{1.00E+100,0.45,15160},
	}
	oldLevels := []info{
		{1500,0.03,0},
		{4500,0.1,105},
		{9000,0.2,555},
		{35000,0.25,1005},
		{55000,0.3,2755},
		{80000,0.35,5505},
		{1.00E+100,0.45,13505},
	}

	oldShould := salary - oldStart - social
	newShould := salary - newStart - social
	if oldShould < 0 {
		oldShould = 0
	}
	if newShould < 0  {
		newShould = 0
	}
	newTax := getTax(newShould,newLevels)
	oldTax := getTax(oldShould,oldLevels)
	data := make(map[string]float64)
	data["oldTax"] = oldTax // 交税额
	data["oldShould"] = oldShould // 需要交税的部分
	data["newTax"] = newTax
	data["newShould"] = newShould
	data["newMore"] = oldTax - newTax
	return data
}

func getTax(should float64, levels []info) (m float64) {
	for _,v := range levels{
		if should <= v.top {
			m = should * v.rates - v.deduction
			break
		}
	}
	return
}
func init() {
}