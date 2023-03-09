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
	"github.com/shirdonl/goDesignPattern/chapter4/observer/actualCombat"
)

func main() {

	//声明商品对象
	bookProduct := actualCombat.NewProduct("《Go语言高级开发与实战》")

	ObserverFirst := &actualCombat.User{Id: "shirdonliao@gmail.com"}
	ObserverSecond := &actualCombat.User{Id: "barry@gmail.com"}

	//通知第一个用户
	bookProduct.Register(ObserverFirst)
	//通知第二个用户
	bookProduct.Register(ObserverSecond)

	//更新库存
	bookProduct.UpdateAvailability()
}

//$ go run main-actual-combat.go
//商品 《Go语言高级开发与实战》 现在上架了～
//发送邮件给用户： shirdonliao@gmail.com ，商品： 《Go语言高级开发与实战》 上架啦～
//发送邮件给用户： barry@gmail.com ，商品： 《Go语言高级开发与实战》 上架啦～
