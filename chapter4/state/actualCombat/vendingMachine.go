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

//自动售货机类
type VendingMachine struct {
	hasProduct       State
	productRequested State
	hasMoney         State
	noProduct        State

	currentState State

	productCount int
	productPrice int
}

func NewVendingMachine(productCount, productPrice int) *VendingMachine {
	v := &VendingMachine{
		productCount: productCount,
		productPrice: productPrice,
	}
	HasProductState := &HasProductState{
		VendingMachine: v,
	}
	ProductRequestedState := &ProductRequestedState{
		VendingMachine: v,
	}
	HasMoneyState := &HasMoneyState{
		VendingMachine: v,
	}
	NoProductState := &NoProductState{
		VendingMachine: v,
	}

	v.SetState(HasProductState)
	v.hasProduct = HasProductState
	v.productRequested = ProductRequestedState
	v.hasMoney = HasMoneyState
	v.noProduct = NoProductState
	return v
}

func (v *VendingMachine) RequestProduct() error {
	return v.currentState.RequestProduct()
}

func (v *VendingMachine) AddProduct(count int) error {
	return v.currentState.AddProduct(count)
}

func (v *VendingMachine) InsertMoney(money int) error {
	return v.currentState.InsertMoney(money)
}

func (v *VendingMachine) DispenseProduct() error {
	return v.currentState.DispenseProduct()
}

func (v *VendingMachine) SetState(s State) {
	v.currentState = s
}

func (v *VendingMachine) IncrementProductCount(count int) {
	fmt.Printf("Adding %d products\n", count)
	v.productCount = v.productCount + count
}
