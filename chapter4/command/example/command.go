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

type Command interface {
	Execute()
}

// 调用者类
type Invoker struct {
	cmds []Command
}

// SetCommand 方法用于设置命令
func (c *Invoker) SetCommand(cmd Command) {
	c.cmds = append(c.cmds, cmd)
}

// ExecuteCommand 方法用于执行命令
func (c *Invoker) ExecuteCommand() {
	for _, cmd := range c.cmds {
		cmd.Execute()
	}
}

// 初始化调用者对象
func NewInvoker() *Invoker {
	return &Invoker{}
}

// 接收者类
type Receiver struct {
}

// 初始化接收者对象
func NewReceiver() *Receiver {
	return &Receiver{}
}

//接收者具体执行操作1
func (f *Receiver) operation1(a string) {
	fmt.Println("operation1:", a)
}

//接收者具体执行操作2
func (f *Receiver) operation2(a, b, c string) {
	fmt.Println("operation2:", a, b, c)
}

// 具体命令1
type Command1 struct {
	name     string
	receiver *Receiver
}

// 初始化Command1对象
func NewCommand1(name string, receiverObj *Receiver) *Command1 {
	return &Command1{
		name:     name,
		receiver: receiverObj,
	}
}

//具体命令1执行操作
func (c *Command1) Execute() {
	c.receiver.operation1(c.name)
}

// 具体命令2
type Command2 struct {
	name     string
	receiver *Receiver
}

// 初始化Command2对象
func NewCommand2(name string, receiverObj *Receiver) *Command2 {
	return &Command2{
		name:     name,
		receiver: receiverObj,
	}
}

//具体命令2执行操作
func (c *Command2) Execute() {
	c.receiver.operation2(c.name, c.name, c.name)
}
