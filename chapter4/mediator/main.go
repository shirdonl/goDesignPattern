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

import "github.com/shirdonl/goDesignPattern/chapter4/mediator/example"

func main() {
	mediator := example.NewMediator()
	mediator.ConcreteColleague2.Talk()
}

//$ go run main.go
//通过中介者谈话
//具体同事1：ConcreteColleague1回复中...
//具体同事2：ConcreteColleague2回复中...
