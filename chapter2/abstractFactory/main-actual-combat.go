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
	"github.com/shirdonl/goDesignPattern/chapter2/abstractFactory/actualCombat"
)

func main() {
	//声明小米工厂
	xiaomiFactory, _ := actualCombat.GetElectronicFactory("Xiaomi")
	//声明联想工厂
	lenovoFactory, _ := actualCombat.GetElectronicFactory("Lenovo")

	//联想工厂生产联想手机
	lenovoPhone := lenovoFactory.MakePhone()
	//联想电脑生产联想电脑
	lenovoComputer := lenovoFactory.MakeComputer()

	//小米工厂生产小米手机
	xiaomiPhone := xiaomiFactory.MakePhone()
	//小米电脑生产小米电脑
	xiaomiComputer := xiaomiFactory.MakeComputer()

	printPhoneDetails(lenovoPhone)
	printComputerDetails(lenovoComputer)

	printPhoneDetails(xiaomiPhone)
	printComputerDetails(xiaomiComputer)
}

func printPhoneDetails(s actualCombat.AbstractPhone) {
	fmt.Printf("Color: %s", s.GetColor())
	fmt.Println()
	fmt.Printf("Size: %d inch", s.GetSize())
	fmt.Println()
}

func printComputerDetails(s actualCombat.AbstractComputer) {
	fmt.Printf("Color: %s", s.GetColor())
	fmt.Println()
	fmt.Printf("Size: %d inch", s.GetSize())
	fmt.Println()
}

//$ go run main-actual-combat.go
//Color: Black
//Size: 5 inch
//Color: White
//Size: 14 inch
//Color: White
//Size: 5 inch
//Color: Black
//Size: 14 inch
