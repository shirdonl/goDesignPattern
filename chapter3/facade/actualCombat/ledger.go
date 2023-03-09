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

//分类帐
type Ledger struct {
}

//生成分类帐条目
func (s *Ledger) MakeEntry(accountID, txnType string, amount int) {
	fmt.Printf("为账户：%s 生成分类帐条目，账目类型为：%s，金额为：%d\n", accountID, txnType, amount)
	return
}
