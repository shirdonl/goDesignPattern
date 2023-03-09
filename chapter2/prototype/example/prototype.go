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

// 原型接口
type Prototype interface {
	GetName() string
	Clone() Prototype
}

// 具体原型类
type ConcretePrototype struct {
	Name string
}

// 返回具体原型的名称
func (p *ConcretePrototype) GetName() string {
	return p.Name
}

// Clone 创建一个ConcretePrototype类的克隆新实例
func (p *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{p.Name}
}
