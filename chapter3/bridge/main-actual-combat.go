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
	"github.com/shirdonl/goDesignPattern/chapter3/bridge/actualCombat"
)

func main() {
	//联想打印机
	lenovoPrinter := &actualCombat.Lenovo{}
	//佳能打印机
	canonPrinter := &actualCombat.Canon{}

	//Mac打印
	macComputer := &actualCombat.Mac{}

	//Mac电脑用SetPrinter()方法设置联想打印机
	macComputer.SetPrinter(lenovoPrinter)
	macComputer.Print()
	fmt.Println()

	//Mac电脑用SetPrinter()方法设置佳能打印机
	macComputer.SetPrinter(canonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &actualCombat.Windows{}
	//Windows电脑用SetPrinter()方法设置联想打印机
	winComputer.SetPrinter(lenovoPrinter)
	winComputer.Print()
	fmt.Println()

	///Windows电脑用SetPrinter()方法设置佳能打印机
	winComputer.SetPrinter(canonPrinter)
	winComputer.Print()
	fmt.Println()
}

//$ go run main-actual-combat.go
//Print request for Mac
//Printing by a Lenovo Printer
//
//Print request for Mac
//Printing by a Canon Printer
//
//Print request for Windows
//Printing by a Lenovo Printer
//
//Print request for Windows
//Printing by a Canon Printer
