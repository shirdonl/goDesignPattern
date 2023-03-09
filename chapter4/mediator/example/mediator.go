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

// 同事类接口
type Colleague interface {
	SetMediator(mediator Mediator)
}

// 具体同事1
type ConcreteColleague1 struct {
	mediator Mediator
}

// 设置中介
func (b *ConcreteColleague1) SetMediator(mediator Mediator) {
	b.mediator = mediator
}

// 执行动作
func (b *ConcreteColleague1) Respond() {
	fmt.Println("具体同事1：ConcreteColleague1回复中...")
	b.mediator.Communicate("ConcreteColleague1")
	return
}

// 具体同事2
type ConcreteColleague2 struct {
	mediator Mediator
}

// 设置中介
func (t *ConcreteColleague2) SetMediator(mediator Mediator) {
	t.mediator = mediator
}

// 通过中介者谈话
func (t *ConcreteColleague2) Talk() {
	fmt.Println("通过中介者谈话")
	t.mediator.Communicate("ConcreteColleague2")
}

// 执行动作
func (t *ConcreteColleague2) Respond() {
	fmt.Println("具体同事2：ConcreteColleague2回复中...")
}

// Mediator 描述了具体同事之间通信的接口
type Mediator interface {
	Communicate(who string)
}

// ConcreateMediator 描述了 ConcreteColleague1 和 ConcreteColleague2 之间的中介
type ConcreteMediator struct {
	ConcreteColleague1
	ConcreteColleague2
}

// NewMediator 创建一个具体中介者 ConcreateMediator
func NewMediator() *ConcreteMediator {
	mediator := &ConcreteMediator{}
	mediator.ConcreteColleague1.SetMediator(mediator)
	mediator.ConcreteColleague2.SetMediator(mediator)
	return mediator
}

// Communicate 在 ConcreteColleague1 和 ConcreteColleague2 之间进行通信
func (m *ConcreteMediator) Communicate(who string) {
	if who == "ConcreteColleague2" {
		m.ConcreteColleague1.Respond()
		return
	} else if who == "ConcreteColleague1" {
		m.ConcreteColleague2.Respond()
		return
	}
}
