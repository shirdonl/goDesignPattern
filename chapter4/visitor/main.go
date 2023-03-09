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

import "github.com/shirdonl/goDesignPattern/chapter4/visitor/example"

func main() {
	//声明具体元素A
	concreteElementA := &example.ConcreteElementA{}
	//调用具体元素A的方法
	concreteElementA.FeatureA()
	//具体元素A接受具体访问者A
	concreteElementA.Accept(&example.ConcreteVisitorA{})

	//声明具体元素B
	concreteElementB := &example.ConcreteElementB{}
	//调用具体元素B的方法
	concreteElementB.FeatureB()
	//具体元素B接受具体访问者B
	concreteElementB.Accept(&example.ConcreteVisitorB{})
}

//$ go run main.go
//具体访问者A ConcreteElementA
//具体访问者B ConcreteElementB
