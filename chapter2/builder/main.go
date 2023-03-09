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
	"github.com/shirdonl/goDesignPattern/chapter2/builder/pkg"
)

func main() {
	//声明Mpv生成器
	MpvBuilder := pkg.GetBuilder("mpv")
	//声明Mpv生成器
	SuvBuilder := pkg.GetBuilder("suv")

	//声明主管
	Director := pkg.NewDirector(MpvBuilder)
	//生产汽车
	mpvCar := Director.BuildCar()

	fmt.Printf("MPV类型引擎: %s\n", mpvCar.EngineType)
	fmt.Printf("MPV类型座椅: %s\n", mpvCar.SeatsType)
	fmt.Printf("MPV类型数量: %d\n", mpvCar.Number)

	//声明主管
	Director.SetBuilder(SuvBuilder)
	//生产汽车
	suvCar := Director.BuildCar()

	fmt.Printf("\nSUV类型引擎: %s\n", suvCar.EngineType)
	fmt.Printf("SUV类型座椅: %s\n", suvCar.SeatsType)
	fmt.Printf("SUV类型数量: %d\n", suvCar.Number)
}
