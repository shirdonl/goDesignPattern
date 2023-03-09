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

import "github.com/shirdonl/goDesignPattern/chapter4/state/example"

func main() {
	context := example.NewContext()
	context.Off()
	context.On()
	context.On()
	context.Off()
}

//$ go run main.go
//Context准备好了
//已经关闭～
//将状态从关闭切换到打开～
//已经打开了～
//将状态从打开切换到关闭～
