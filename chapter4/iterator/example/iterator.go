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

package example

import (
	"fmt"
	"time"
)

// 迭代器接口
type Iterator interface {
	// 返回是否存在另一个下一个元素
	HasMore() bool

	// 递增迭代器，用于指向下一个元素
	GetNext()
}

//集合接口
type Collection interface {
	CreateIterator() Iterator
}

//具体集合
type ConcreteCollection struct {
}

//初始化具体集合对象，用于创建具体迭代器对象
func (u *ConcreteCollection) CreateIterator() Iterator {
	return &ConcreteIterator{
		IterationState: true,
	}
}

// 具体迭代器
type ConcreteIterator struct {
	IterationState bool
}

//具体迭代器的方法
func (i *ConcreteIterator) HasMore() bool {
	if i.IterationState == true {
		return true
	} else {
		return false
	}
}

//具体迭代器的方法，用于递增迭代器以指向下一个元素
func (i *ConcreteIterator) GetNext() {
	if i.HasMore() {
		time.Sleep(1 * time.Second)
		fmt.Println("GetNext")
	}
}
