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
	"github.com/shirdonl/goDesignPattern/chapter3/facade/example"
)

func main() {
	fa := example.NewFacade()
	fa.MethodA()
	fa.MethodB()

	sub := example.NewSubSystemA()
	sub.MethodOne()
	sub.MethodTwo()
}

//$ go run main.go
//SubSystemA - MethodThree
//SubSystemB - MethodOne
//SubSystemA - MethodFour
//SubSystemA - MethodFour
//SubSystemB - MethodTwo
//SubSystemB - MethodOne
//SubSystemB - MethodTwo
