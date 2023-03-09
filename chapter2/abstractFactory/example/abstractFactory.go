//++++++++++++++++++++++++++++++++++++++++
// 《Go语言设计模式》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goDesignPattern
// 仓库地址：https://github.com/shirdonl/goDesignPattern
// 交流咨询，请关注公众号"源码大数据"
//++++++++++++++++++++++++++++++++++++++++

package example

import (
	"fmt"
)

// 抽象产品接口
type AbstractProduct interface {
	GetName()
}

//具体产品
type ConcreteProduct struct {
}

//具体产品的方法
func (c *ConcreteProduct) GetName() {
	fmt.Println("具体产品 ConcreteProduct")
}

// 抽象工厂接口
type AbstractFactory interface {
	CreateProduct() AbstractProduct
}

//具体工厂
type ConcreteFactory struct {
}

// 初始化具体工厂对象
func NewConcreteFactory() ConcreteFactory {
	return ConcreteFactory{}
}

//具体工厂创建具体产品
func (s *ConcreteFactory) CreateProduct() ConcreteProduct {
	return ConcreteProduct{}
}
