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

import "github.com/shirdonl/goDesignPattern/chapter4/command/example"

func main() {
	//创建接收者
	receiver := example.NewReceiver()
	cc := example.NewInvoker()
	//创建具体命令对象，如有需要可将其关联至接收者
	cmd1 := example.NewCommand1("commandA", receiver)
	cmd2 := example.NewCommand2("commandB", receiver)
	cc.SetCommand(cmd1)
	cc.SetCommand(cmd2)
	//执行命令
	cc.ExecuteCommand()
}

//operation1: commandA
//operation2: commandB commandB commandB
