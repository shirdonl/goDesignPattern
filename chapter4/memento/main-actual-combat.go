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
	"github.com/shirdonl/goDesignPattern/chapter4/memento/actualCombat"
)

func main() {

	//声明负责人对象
	Caretaker := &actualCombat.Caretaker{
		MementoArray: make([]*actualCombat.Memento, 0),
	}

	//声明原发器对象
	Originator := new(actualCombat.Originator)
	Originator.State = "One"

	fmt.Printf("Originator 当前状态: %s\n", Originator.GetState())
	//添加备忘录
	Caretaker.AddMemento(Originator.CreateMemento())

	Originator.SetState("Two")
	fmt.Printf("Originator 当前状态: %s\n", Originator.GetState())
	//添加备忘录
	Caretaker.AddMemento(Originator.CreateMemento())

	Originator.SetState("Three")
	fmt.Printf("Originator 当前状态: %s\n", Originator.GetState())
	//添加备忘录
	Caretaker.AddMemento(Originator.CreateMemento())

	//恢复原发器对象的状态
	Originator.RestoreMemento(Caretaker.GetMemento(1))
	fmt.Printf("恢复到状态: %s\n", Originator.GetState())
	//恢复原发器对象的状态
	Originator.RestoreMemento(Caretaker.GetMemento(0))
	fmt.Printf("恢复到状态: %s\n", Originator.GetState())
}
