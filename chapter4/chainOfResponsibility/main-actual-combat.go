//++++++++++++++++++++++++++++++++++++++++
// 《Go语言设计模式》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 作者知乎：https://www.zhihu.com/people/shirdonl
// 仓库地址：https://gitee.com/shirdonl/goDesignPattern
// 仓库地址：https://github.com/shirdonl/goDesignPattern
// 交流咨询，请关注公众号"源码大数据"
//++++++++++++++++++++++++++++++++++++++++

package main

import (
	"github.com/shirdonl/goDesignPattern/chapter4/chainOfResponsibility/actualCombat"
)

func main() {

	cashier := &actualCombat.Cashier{}

	//设置下一个医务部门
	medical := &actualCombat.Drugstore{}
	medical.SetNext(cashier)

	//设置下一个医务部门
	doctor := &actualCombat.Clinic{}
	doctor.SetNext(medical)

	//设置下一个医务部门
	reception := &actualCombat.Reception{}
	reception.SetNext(doctor)

	patient := &actualCombat.Patient{Name: "Jack"}
	//设置病人
	reception.Execute(patient)
}

//$ go run main-actual-combat.go
//正在接待登记病人
//医生正在检查病人
//正在给病人用药
//收银员从病人那里收钱
