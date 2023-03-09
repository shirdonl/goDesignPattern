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

// 上下文定义了可以打开和关闭的状态
type Context struct {
	state State
}

// 初始化上下文
func NewContext() *Context {
	fmt.Println("Context准备好了")
	return &Context{NewConcreteStateB()}
}

// ChangeState 设置上下文的当前状态
func (c *Context) ChangeState(s State) {
	c.state = s
}

// 按下打开按钮
func (c *Context) On() {
	c.state.On(c)
}

// 按下关闭按钮
func (c *Context) Off() {
	c.state.Off(c)
}

// 描述了上下文的内部状态
type State interface {
	On(m *Context)
	Off(m *Context)
}

//ConcreteStateA 描述了打开按钮的状态
type ConcreteStateA struct {
	context Context
}

// 创建一个新的具体状态A
func NewConcreteStateA() State {
	return &ConcreteStateA{}
}

// 具体状态A的打开方法
func (cs *ConcreteStateA) On(m *Context) {
	fmt.Println("已经打开了～")
}

// 将状态从打开切换到关闭
func (cs *ConcreteStateA) Off(m *Context) {
	fmt.Println("将状态从打开切换到关闭～")
	m.ChangeState(NewConcreteStateB())
}

// ConcreteStateB 描述关闭按钮状态
type ConcreteStateB struct {
	context Context
}

// NewConcreteStateB 创建一个新的具体状态B
func NewConcreteStateB() State {
	return &ConcreteStateB{}
}

// 将状态从关闭切换到打开
func (cs *ConcreteStateB) On(m *Context) {
	fmt.Println("将状态从关闭切换到打开～")
	m.ChangeState(NewConcreteStateA())
}

// 具体状态B的关闭方法
func (cs *ConcreteStateB) Off(m *Context) {
	fmt.Println("已经关闭～")
}
