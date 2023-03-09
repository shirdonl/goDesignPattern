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
	"github.com/shirdonl/goDesignPattern/chapter4/strategy/example"
)

func main() {
	strategyB := example.NewStrategyB()
	context := example.NewContext()
	context.SetStrategy(strategyB)
	strategyA := example.NewStrategyA()
	context.SetStrategy(strategyA)
	context.Execute()
}

//$ go run main.go
//执行策略 A
