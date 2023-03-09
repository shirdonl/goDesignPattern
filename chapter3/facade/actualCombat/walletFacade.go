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

//定义钱包的外观类
type WalletFacade struct {
	Account          *Account
	Wallet           *Wallet
	VerificationCode *VerificationCode
	Notification     *Notification
	Ledger           *Ledger
}

//创建钱包的外观类
func NewWalletFacade(accountID string, code int) *WalletFacade {
	WalletFacacde := &WalletFacade{
		Account:          NewAccount(accountID),
		VerificationCode: NewVerificationCode(code),
		Wallet:           NewWallet(),
		Notification:     &Notification{},
		Ledger:           &Ledger{},
	}
	return WalletFacacde
}

//添加钱到钱包
func (w *WalletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("添加钱到钱包")
	//1.检查账户
	err := w.Account.CheckAccount(accountID)
	if err != nil {
		return err
	}
	//2.检查验证码
	err = w.VerificationCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	//3.添加金额
	w.Wallet.AddBalance(amount)
	//4.发送信用通知
	w.Notification.SendWalletCreditNotification()
	w.Ledger.MakeEntry(accountID, "credit", amount)
	return nil
}

//从钱包里扣款
func (w *WalletFacade) DeductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("从钱包里扣款")
	//1.检查账户
	err := w.Account.CheckAccount(accountID)
	if err != nil {
		return err
	}

	//2.检查验证码
	err = w.VerificationCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	//3.借款金额
	err = w.Wallet.DebitBalance(amount)
	if err != nil {
		return err
	}
	//4.发送借款通知
	w.Notification.SendWalletDebitNotification()
	w.Ledger.MakeEntry(accountID, "credit", amount)
	return nil
}
