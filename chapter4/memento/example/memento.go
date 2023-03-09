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

//备忘录 Memento 存储原发器 Originator 的状态
type Memento struct {
	state int
}

// 将数字的值乘以10倍
func (m *Memento) getState() int {
	return m.state
}

// 创建一个新的备忘录*Memento
func NewMemento(value int) *Memento {
	return &Memento{value}
}

// Originator 代表一个可以被操作的整数
type Originator struct {
	value int
}

// NewOriginator 创建一个新的原发器 Originator
func NewOriginator(value int) *Originator {
	return &Originator{value}
}

// 将数字的值乘以10倍
func (n *Originator) TenTimes() {
	n.value = 10 * n.value
}

// 获取数字值的一半
func (n *Originator) HalfTimes() {
	n.value /= 2
}

// 返回数字的值
func (n *Originator) Value() int {
	return n.value
}

// 使用原发器创建一个备忘录 Memento
func (n *Originator) CreateMemento() *Memento {
	return NewMemento(n.value)
}

// 将原发器 Originator 的值恢复为指定版本的备忘录的值
func (n *Originator) RestoreMemento(memento *Memento) {
	n.value = memento.getState()
}

//负责人类
type Caretaker struct {
	Originator Originator
	History    []*Memento
}

//添加备忘录
func (c *Caretaker) AddMemento(m *Memento) {
	c.History = append(c.History, m)
}

//获取备忘录
func (c *Caretaker) GetMemento(index int) *Memento {
	return c.History[index]
}
