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

import "fmt"

import (
	"github.com/shirdonl/goDesignPattern/chapter4/visitor/actualCombat"
)

func main() {
	//声明边长为1的正方形
	Square := &actualCombat.Square{Side: 1}
	//声明半径为6的圆形
	Circle := &actualCombat.Circle{Radius: 6}
	//声明长为8，宽为6的矩形
	Rectangle := &actualCombat.Rectangle{L: 8, B: 6}

	//声明周长计算器
	PerimeterCalculator := &actualCombat.PerimeterCalculator{}

	//正方形计算周长
	Square.Accept(PerimeterCalculator)
	//圆形计算周长
	Circle.Accept(PerimeterCalculator)
	//矩形计算周长
	Rectangle.Accept(PerimeterCalculator)
	Square.GetType()

	fmt.Println()
	//声明中点坐标计算器
	MiddleCoordinates := &actualCombat.MiddleCoordinates{}
	//获取正方形中点坐标
	Square.Accept(MiddleCoordinates)
	//获取圆形中点坐标
	Circle.Accept(MiddleCoordinates)
	//获取矩形中点坐标
	Rectangle.Accept(MiddleCoordinates)
}

//$ go run main-actual-combat.go
//计算正方形的周长：4.000000
//计算圆形的周长：37.680000
//计算矩形的周长：28.000000
//
//计算正方形的中点坐标
//计算圆形的中点坐标
//计算矩形的中点坐标
