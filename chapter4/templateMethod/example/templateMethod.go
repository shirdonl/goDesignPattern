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
)

// 抽象类接口
type AbstractClassInterface interface {
	Step1()
	Step2()
	Step3()
}

// 抽象类
type AbstractClass struct {
	AbstractClassInterface
}

// 初始化抽象类对象
func NewAbstractClass(aci AbstractClassInterface) *AbstractClass {
	return &AbstractClass{aci}
}

// 模版方法
func (cc *AbstractClass) TemplateMethod() {
	cc.Step1()
	cc.Step2()
	cc.Step3()
}

// 具体类A
type ConcreteClassA struct {
}

// 具体类A的方法1
func (cc *ConcreteClassA) Step1() {
	fmt.Println("ConcreteClassA Step1")
}

// 具体类A的方法2
func (cc *ConcreteClassA) Step2() {
	fmt.Println("ConcreteClassA Step2")
}

// 具体类A的方法3
func (cc *ConcreteClassA) Step3() {
	fmt.Println("ConcreteClassA Step3")
}

// 具体类B
type ConcreteClassB struct {
}

// 具体类B的方法1
func (cc *ConcreteClassB) Step1() {
	fmt.Println("ConcreteClassB Step1")
}

// 具体类B的方法2
func (cc *ConcreteClassB) Step2() {
	fmt.Println("ConcreteClassB Step2")
}

// 具体类B的方法3
func (cc *ConcreteClassB) Step3() {
	fmt.Println("ConcreteClassB Step3")
}
