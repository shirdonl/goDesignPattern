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

type HasMoneyState struct {
	VendingMachine *VendingMachine
}

func (i *HasMoneyState) RequestProduct() error {
	return fmt.Errorf("Product dispense in progress")
}

func (i *HasMoneyState) AddProduct(count int) error {
	return fmt.Errorf("Product dispense in progress")
}

func (i *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("Product out of stock")
}
func (i *HasMoneyState) DispenseProduct() error {
	fmt.Println("Dispensing Product")
	i.VendingMachine.productCount = i.VendingMachine.productCount - 1
	if i.VendingMachine.productCount == 0 {
		i.VendingMachine.SetState(i.VendingMachine.noProduct)
	} else {
		i.VendingMachine.SetState(i.VendingMachine.hasProduct)
	}
	return nil
}
