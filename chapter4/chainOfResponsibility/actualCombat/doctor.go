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

type Clinic struct {
	next department
}

func (d *Clinic) Execute(p *Patient) {
	if p.ClinicCheckUpDone {
		fmt.Println("医生已经检查过了")
		d.next.Execute(p)
		return
	}
	fmt.Println("医生正在检查病人")
	p.ClinicCheckUpDone = true
	d.next.Execute(p)
}

func (d *Clinic) SetNext(next department) {
	d.next = next
}
