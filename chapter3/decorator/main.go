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

import "github.com/shirdonl/goDesignPattern/chapter3/decorator/example"

func main() {
	concreteComponent := &example.ConcreteComponent{}
	decoratorA := &example.DecoratorA{}
	decoratorB := &example.DecoratorB{}
	decoratorA.SetComponent(concreteComponent)
	decoratorB.SetComponent(decoratorA)
	decoratorB.Operation()
}

//$ go run main.go
//具体的对象开始操作...
//装饰A扩展的方法~
//装饰B扩展的方法~
