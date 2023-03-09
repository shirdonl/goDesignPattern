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
	"github.com/shirdonl/goDesignPattern/chapter3/facade/actualCombat"
	"log"
)

func main() {
	fmt.Println()
	//实例化外观模式
	WalletFacade := actualCombat.NewWalletFacade("barry", 1688)
	fmt.Println()

	//添加16元到钱包
	err := WalletFacade.AddMoneyToWallet("barry", 1688, 16)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	//从钱包取出5元
	err = WalletFacade.DeductMoneyFromWallet("barry", 1688, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}

//$ go run main-actual-combat.go
//
//添加钱到钱包
//账户验证通过～
//验证通过～
//添加钱包金额成功～
//发送钱包信用通知...
//为账户：barry 生成分类帐条目，账目类型为：credit，金额为：16
//
//从钱包里扣款
//账户验证通过～
//验证通过～
//钱包金额足够～
//发送钱包借款通知...
//为账户：barry 生成分类帐条目，账目类型为：credit，金额为：5
