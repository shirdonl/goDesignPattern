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

type Cashier struct {
	next department
}

func (c *Cashier) Execute(p *Patient) {
	if p.PaymentDone {
		fmt.Println("支付完成")
	}
	fmt.Println("收银员从病人那里收钱")
}

func (c *Cashier) SetNext(next department) {
	c.next = next
}
