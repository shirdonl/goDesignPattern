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

import "fmt"

type Drugstore struct {
	next department
}

func (m *Drugstore) Execute(p *Patient) {
	if p.MedicineDone {
		fmt.Println("药品已经给病人")
		m.next.Execute(p)
		return
	}
	fmt.Println("正在给病人用药")
	p.MedicineDone = true
	m.next.Execute(p)
}

func (m *Drugstore) SetNext(next department) {
	m.next = next
}
