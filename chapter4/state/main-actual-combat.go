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
	"github.com/shirdonl/goDesignPattern/chapter4/state/actualCombat"
	"log"
)

func main() {
	//声明自动售货机
	VendingMachine := actualCombat.NewVendingMachine(1, 10)

	//请求商品
	err := VendingMachine.RequestProduct()
	if err != nil {
		log.Fatalf(err.Error())
	}

	//向自动售货机放入10元钱
	err = VendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	//向自动售货机分配商品
	err = VendingMachine.DispenseProduct()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	//向自动售货机添加2个商品
	err = VendingMachine.AddProduct(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()
}

//$ go run main-actual-combat.go
//Product requestd
//Money entered is ok
//Dispensing Product
//
//Adding 2 products
