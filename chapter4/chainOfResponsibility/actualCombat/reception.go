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

type Reception struct {
	next department
}

func (r *Reception) Execute(p *Patient) {
	if p.RegistrationDone {
		fmt.Println("已完成患者登记")
		r.next.Execute(p)
		return
	}
	fmt.Println("正在接待登记病人")
	p.RegistrationDone = true
	r.next.Execute(p)
}

func (r *Reception) SetNext(next department) {
	r.next = next
}
