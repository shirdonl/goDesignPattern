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

type ProductRequestedState struct {
	VendingMachine *VendingMachine
}

func (i *ProductRequestedState) RequestProduct() error {
	return fmt.Errorf("Product already requested")
}

func (i *ProductRequestedState) AddProduct(count int) error {
	return fmt.Errorf("Product Dispense in progress")
}

func (i *ProductRequestedState) InsertMoney(money int) error {
	if money < i.VendingMachine.productPrice {
		fmt.Errorf("Inserted money is less. Please insert %d", i.VendingMachine.productPrice)
	}
	fmt.Println("Money entered is ok")
	i.VendingMachine.SetState(i.VendingMachine.hasMoney)
	return nil
}
func (i *ProductRequestedState) DispenseProduct() error {
	return fmt.Errorf("Please insert money first")
}
