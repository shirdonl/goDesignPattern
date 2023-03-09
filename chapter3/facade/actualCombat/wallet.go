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

//钱包
type Wallet struct {
	balance int
}

//创建钱包
func NewWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}

//添加金额
func (w *Wallet) AddBalance(amount int) {
	w.balance += amount
	fmt.Println("添加钱包金额成功～")
	return
}

//借款金额
func (w *Wallet) DebitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("%s", "金额无效～")
	}
	fmt.Println("钱包金额足够～")
	w.balance = w.balance - amount
	return nil
}
