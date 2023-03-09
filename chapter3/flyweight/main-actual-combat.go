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
	"github.com/shirdonl/goDesignPattern/chapter3/flyweight/actualCombat"
)

func main() {
	game := actualCombat.NewGame{}

	//创建红队
	game.AddBlueTeam(actualCombat.BlueTeamDressType)
	game.AddBlueTeam(actualCombat.BlueTeamDressType)
	game.AddBlueTeam(actualCombat.BlueTeamDressType)
	game.AddBlueTeam(actualCombat.BlueTeamDressType)

	//创建蓝队
	game.AddRedTeam(actualCombat.RedTeamDressType)
	game.AddRedTeam(actualCombat.RedTeamDressType)
	game.AddRedTeam(actualCombat.RedTeamDressType)

	DressFactoryInstance := actualCombat.GetDressFactorySingleInstance()

	for DressType, Dress := range DressFactoryInstance.DressMap {
		fmt.Printf("服装类型: %s\n服装颜色: %s\n", DressType, Dress.GetColor())
	}
}

//$ go run main-actual-combat.go
//服装类型: Blue Dress
//服装颜色: blue
//服装类型: Red Dress
//服装颜色: red
