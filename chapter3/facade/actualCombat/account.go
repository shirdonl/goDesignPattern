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

//账户
type Account struct {
	name string
}

//创建账户
func NewAccount(accountName string) *Account {
	return &Account{
		name: accountName,
	}
}

//检查账户
func (a *Account) CheckAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("%s", "账户名不正确～")
	}
	fmt.Println("账户验证通过～")
	return nil
}
