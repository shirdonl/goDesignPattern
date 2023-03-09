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
	"github.com/shirdonl/goDesignPattern/chapter4/iterator/actualCombat"
)

func main() {

	//声明用户对象User1
	user1 := &actualCombat.User{
		Name: "Jack",
		Age:  30,
	}
	//声明用户对象User2
	user2 := &actualCombat.User{
		Name: "Barry",
		Age:  20,
	}

	//声明具体集合对象
	userCollection := &actualCombat.UserCollection{
		Users: []*actualCombat.User{user1, user2},
	}

	//声明具体迭代器对象
	iterator := userCollection.CreateIterator()

	//执行具体方法
	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}

//$ go run main-actual-combat.go
//User is &{Name:Jack Age:30}
//User is &{Name:Barry Age:20}
