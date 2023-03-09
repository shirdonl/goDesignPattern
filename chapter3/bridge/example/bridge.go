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

//实现接口
type Implementor interface {
	Implementation(str string)
}

//具体实现
type ConcreteImplementor struct{}

func (*ConcreteImplementor) Implementation(str string) {
	fmt.Printf("打印信息：[%v]", str)
}

//初始化具体实现对象
func NewConcreteImplementor() *ConcreteImplementor {
	return &ConcreteImplementor{}
}

//抽象接口
type Abstraction interface {
	Execute(str string)
}

//扩充抽象
type RefinedAbstraction struct {
	method Implementor
}

//扩充抽象方法
func (c *RefinedAbstraction) Execute(str string) {
	c.method.Implementation(str)
}

//初始化扩充抽象对象
func NewRefinedAbstraction(im Implementor) *RefinedAbstraction {
	return &RefinedAbstraction{method: im}
}
