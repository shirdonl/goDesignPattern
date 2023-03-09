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

type NoProductState struct {
	VendingMachine *VendingMachine
}

func (i *NoProductState) RequestProduct() error {
	return fmt.Errorf("Product out of stock")
}

func (i *NoProductState) AddProduct(count int) error {
	i.VendingMachine.IncrementProductCount(count)
	i.VendingMachine.SetState(i.VendingMachine.hasProduct)
	return nil
}

func (i *NoProductState) InsertMoney(money int) error {
	return fmt.Errorf("Product out of stock")
}
func (i *NoProductState) DispenseProduct() error {
	return fmt.Errorf("Product out of stock")
}
