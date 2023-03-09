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

package pkg

import "fmt"

//通知
type Notification struct {
}

//发送信用通知
func (n *Notification) SendWalletCreditNotification() {
	fmt.Println("发送钱包信用通知...")
}

//发送借款通知
func (n *Notification) SendWalletDebitNotification() {
	fmt.Println("发送钱包借款通知...")
}

