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
	"github.com/shirdonl/goDesignPattern/chapter3/decorator/actualCombat"
)

func main() {
	//具体零件
	phone := &actualCombat.BaseParts{}
	fmt.Printf("基础零件的价格为：%f\n", phone.GetPrice())

	//定义添加IPhone手机
	iPhone := &actualCombat.IPhone{}
	iPhone.SetComponent(phone)
	fmt.Printf("苹果的价格为：%f\n", iPhone.GetPrice())

	//定义添加Xiaomi手机
	xiaomi := &actualCombat.Xiaomi{}
	xiaomi.SetComponent(phone)
	fmt.Printf("小米的价格为：%f\n", xiaomi.GetPrice())
}

//$ go run main-actual-combat.go
//基础零件的价格为：2000.000000
//苹果的价格为：8000.000000
//小米的价格为：3000.000000
