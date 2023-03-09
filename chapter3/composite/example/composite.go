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

// 组件接口
type Component interface {
	Execute()
}

// 叶节点，用于描述层次结构中的原始叶节点对象
type Leaf struct {
	value int
}

// 创建一个新的叶节点对象
func NewLeaf(value int) *Leaf {
	return &Leaf{value}
}

// 打印叶节点对象的值
func (l *Leaf) Execute() {
	fmt.Printf("%v  ", l.value)
}

// 组件的组合
type Composite struct {
	children []Component
}

// 创建一个新的组合对象
func NewComposite() *Composite {
	return &Composite{make([]Component, 0)}
}

// 将一个新组件添加到组合中
func (c *Composite) Add(component Component) {
	c.children = append(c.children, component)
}

// 遍历复合子对象
func (c *Composite) Execute() {
	for i := 0; i < len(c.children); i++ {
		c.children[i].Execute()
	}
}
