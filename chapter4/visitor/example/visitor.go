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

// 元素定义了一个接受访问者的接口
type Element interface {
	Accept(v Visitor)
}

// 具体元素A
type ConcreteElementA struct {
}

// 具体元素A的方法
func (t *ConcreteElementA) FeatureA() string {
	return "ConcreteElementA"
}

// Accept 接受访问者
func (t *ConcreteElementA) Accept(v Visitor) {
	v.VisitConcreteElementA(t)
}

// 具体元素B
type ConcreteElementB struct {
}

// 具体元素B的方法
func (t *ConcreteElementB) FeatureB() string {
	return "ConcreteElementB"
}

// Accept 接受访问者
func (t *ConcreteElementB) Accept(v Visitor) {
	v.VisitConcreteElementB(t)
}

// 访问者接口
type Visitor interface {
	VisitConcreteElementA(e *ConcreteElementA)
	VisitConcreteElementB(e *ConcreteElementB)
}

// 具体访问者A
type ConcreteVisitorA struct {
}

// 具体访问者A的方法
func (v *ConcreteVisitorA) VisitConcreteElementA(e *ConcreteElementA) {
	fmt.Printf("具体访问者A %v\n", e.FeatureA())
}

// 具体访问者A的方法
func (v *ConcreteVisitorA) VisitConcreteElementB(e *ConcreteElementB) {
	fmt.Printf("具体访问者A %v\n", e.FeatureB())
}

// 具体访问者B
type ConcreteVisitorB struct {
}

// 具体访问者B的方法
func (v *ConcreteVisitorB) VisitConcreteElementA(e *ConcreteElementA) {
	fmt.Printf("具体访问者B %v\n", e.FeatureA())
}

// 具体访问者B的方法
func (v *ConcreteVisitorB) VisitConcreteElementB(e *ConcreteElementB) {
	fmt.Printf("具体访问者B %v\n", e.FeatureB())
}
