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
	"github.com/shirdonl/goDesignPattern/chapter4/strategy/actualCombat"
)

func main() {
	//声明Lfu对象
	lfu := &actualCombat.Lfu{}
	//初始化缓存对象
	cache := actualCombat.InitCache(lfu)

	//添加缓存
	cache.Add("one", "1")
	cache.Add("two", "2")

	cache.Add("three", "3")

	//声明Lru对象
	lru := &actualCombat.Lru{}
	//设置lru算法类型
	cache.SetAlgorithmType(lru)

	//添加缓存
	cache.Add("four", "4")

	//声明Fifo对象
	fifo := &actualCombat.Fifo{}
	//设置Fifo算法类型
	cache.SetAlgorithmType(fifo)

	//添加缓存
	cache.Add("five", "5")
}

//$ go run main-actual-combat.go
//Deleting by lfu strategy
//Deleting by lru strategy
//Deleting by fifo strategy
