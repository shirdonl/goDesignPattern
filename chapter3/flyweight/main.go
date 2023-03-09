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
	"fmt"
	"github.com/shirdonl/goDesignPattern/chapter3/flyweight/example"
)

func main() {
	factory := example.NewFlyweightFactory()
	flyweight1 := factory.GetFlyweight("Barry")
	flyweight2 := factory.GetFlyweight("Shirdon")

	fmt.Println(flyweight1.Operation("ok"))
	fmt.Println(flyweight2.Operation("good"))
}

//$ go run main.go
//Barry
//ok
//Shirdon
//good
