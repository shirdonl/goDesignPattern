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

//策略接口
type Strategy interface {
	Execute()
}

// 具体策略 A
type strategyA struct {
}

// 具体策略 A 的方法
func (s *strategyA) Execute() {
	fmt.Println("执行策略 A")
}

//具体策略B
type strategyB struct {
}

//具体策略B的方法
func (s *strategyB) Execute() {
	fmt.Println("执行策略 B")
}

//上下文
type Context struct {
	strategy Strategy
}

//设置上下文执行的策略
func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

//上下文的方法
func (c *Context) Execute() {
	c.strategy.Execute()
}

//创建策略 A 的新对象
func NewStrategyA() Strategy {
	return &strategyA{}
}

//创建策略 B 的新对象
func NewStrategyB() Strategy {
	return &strategyB{}
}

//创建一个新的上下文对象
func NewContext() *Context {
	return &Context{}
}
