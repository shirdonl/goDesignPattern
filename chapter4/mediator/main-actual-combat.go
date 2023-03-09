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
	"github.com/shirdonl/goDesignPattern/chapter4/mediator/actualCombat"
)

func main() {
	//声明具体中介者
	stationManager := actualCombat.NewStationManger()

	//声明客运火车
	passengerTrain := &actualCombat.PassengerTrain{
		Mediator: stationManager,
	}
	//声明货运火车
	freightTrain := &actualCombat.FreightTrain{
		Mediator: stationManager,
	}

	passengerTrain.Arrive()
	freightTrain.Arrive()
	passengerTrain.Depart()
}

//$ go run main-actual-combat.go
//PassengerTrain: Arrived
//FreightTrain: Arrival blocked, waiting
//PassengerTrain: Leaving
//FreightTrain: Arrival permitted
//FreightTrain: Arrived
