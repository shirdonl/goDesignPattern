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
	"github.com/shirdonl/goDesignPattern/chapter5/specification/actualCombat"
)

func main() {
	//声明发票数据到期规格
	overDue := actualCombat.NewOverDueSpecification()
	//声明发票通知发送规格
	noticeSent := actualCombat.NewNoticeSentSpecification()
	//收款机构是否收到发票通知
	inCollection := actualCombat.NewInCollectionSpecification()

	sendToCollection := overDue.And(noticeSent).And(inCollection.Not())

	object := actualCombat.Invoice{
		Day:    32,    // >= 30
		Notice: 6,     // >= 3
		IsSent: false, // false
	}

	// 检查规格
	result := sendToCollection.IsSatisfiedBy(object)
	fmt.Println(result)
}

//$ go run main.go
//true
