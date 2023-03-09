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

package actualCombat

//原发器类
type Originator struct {
	State string
}

//创建备忘录
func (e *Originator) CreateMemento() *Memento {
	return &Memento{State: e.State}
}

//恢复原发器对象的状态
func (e *Originator) RestoreMemento(m *Memento) {
	e.State = m.GetSavedState()
}

//设置状态
func (e *Originator) SetState(State string) {
	e.State = State
}

//获取状态
func (e *Originator) GetState() string {
	return e.State
}
