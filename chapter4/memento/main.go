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
	"github.com/shirdonl/goDesignPattern/chapter4/memento/example"
)

func main() {
	//声明负责人对象
	Caretaker := &example.Caretaker{
		History: make([]*example.Memento, 0),
	}

	//声明原发器对象
	n := example.NewOriginator(100)

	//添加备忘录
	Caretaker.AddMemento(n.CreateMemento())
	n.TenTimes()
	fmt.Printf("Originator 当前的值: %d\n", n.Value())

	//添加备忘录
	Caretaker.AddMemento(n.CreateMemento())
	n.TenTimes()
	fmt.Printf("Originator 当前的值: %d\n", n.Value())

	//恢复原发器对象的值
	n.RestoreMemento(Caretaker.GetMemento(0))
	fmt.Printf("恢复备忘录后 Originator 当前的值: %d\n", n.Value())
}

//$ go run main.go
//Originator 当前的值: 1000
//Originator 当前的值: 10000
//恢复备忘录后 Originator 当前的值: 100
