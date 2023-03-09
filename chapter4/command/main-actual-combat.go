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
	"github.com/shirdonl/goDesignPattern/chapter4/command/actualCombat"
)

func main() {
	//初始化具体接收者对象
	Light := &actualCombat.Light{}

	//发送打开命令
	onCommand := &actualCombat.OnCommand{
		Device: Light,
	}

	//发送关闭命令
	offCommand := &actualCombat.OffCommand{
		Device: Light,
	}

	//接收打开命令
	onButton := &actualCombat.Button{
		Command: onCommand,
	}
	//按打开命令键
	onButton.Press()

	//接收关闭命令
	offButton := &actualCombat.Button{
		Command: offCommand,
	}
	//按关闭命令键
	offButton.Press()
}

//$ go run main-actual-combat.go
//打开灯...
//关闭灯...
