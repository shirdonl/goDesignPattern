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

import "fmt"

// 工厂接口
type Factory interface {
	FactoryMethod(owner string) Product
}

// 具体工厂
type ConcreteFactory struct {
}

// 具体工厂的工厂方法
func (cf *ConcreteFactory) FactoryMethod(owner string) Product {
	p := &ConcreteProduct{}
	return p
}

// 产品
type Product interface {
	Use()
}

//具体产品
type ConcreteProduct struct {
}

//具体产品的方法
func (p *ConcreteProduct) Use() {
	fmt.Println("This is a concrete product")
}
