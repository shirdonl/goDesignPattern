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
	"github.com/shirdonl/goDesignPattern/chapter2/factory/pkg"
)

func main() {
	ANTA, _ := pkg.MakeClothes("ANTA")
	PEAK, _ := pkg.MakeClothes("PEAK")

	printDetails(ANTA)
	printDetails(PEAK)
}

func printDetails(c pkg.IClothes) {
	fmt.Printf("Clothes: %s", c.GetName())
	fmt.Println()
	fmt.Printf("Size: %d", c.GetSize())
	fmt.Println()
}

//Clothes: ANTA clothes
//Size: 4
//Clothes: PEAK clothes
//Size: 1
