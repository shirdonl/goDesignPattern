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

type HasProductState struct {
	VendingMachine *VendingMachine
}

func (i *HasProductState) RequestProduct() error {
	if i.VendingMachine.productCount == 0 {
		i.VendingMachine.SetState(i.VendingMachine.noProduct)
		return fmt.Errorf("No product present")
	}
	fmt.Printf("Product requestd\n")
	i.VendingMachine.SetState(i.VendingMachine.productRequested)
	return nil
}

func (i *HasProductState) AddProduct(count int) error {
	fmt.Printf("%d products added\n", count)
	i.VendingMachine.IncrementProductCount(count)
	return nil
}

func (i *HasProductState) InsertMoney(money int) error {
	return fmt.Errorf("Please select product first")
}
func (i *HasProductState) DispenseProduct() error {
	return fmt.Errorf("Please select product first")
}
