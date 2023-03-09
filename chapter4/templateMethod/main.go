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

import "github.com/shirdonl/goDesignPattern/chapter4/templateMethod/example"

func main() {
	concreteClassA := example.NewAbstractClass(&example.ConcreteClassA{})
	concreteClassA.TemplateMethod()
	concreteClassB := example.NewAbstractClass(&example.ConcreteClassB{})
	concreteClassB.TemplateMethod()
}

//$ go run main.go
//ConcreteClassA Step1
//ConcreteClassA Step2
//ConcreteClassA Step3
//ConcreteClassB Step1
//ConcreteClassB Step2
//ConcreteClassB Step3
