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
	"github.com/shirdonl/goDesignPattern/chapter4/iterator/example"
)

func main() {

	//声明具体集合对象
	concreteCollection := &example.ConcreteCollection{}

	//声明具体迭代器对象
	iterator := concreteCollection.CreateIterator()

	//执行具体方法
	for iterator.HasMore() {
		iterator.GetNext()
	}
}
