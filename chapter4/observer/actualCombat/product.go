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

//商品类
type Product struct {
	ObserverList []Observer
	name         string
	inStock      bool
}

//新增商品
func NewProduct(name string) *Product {
	return &Product{
		name: name,
	}
}

//更新库存状态
func (i *Product) UpdateAvailability() {
	fmt.Printf("商品 %s 现在上架了～\n", i.name)
	i.inStock = true
	i.NotifyAll()
}

//注册到观察者列表
func (i *Product) Register(o Observer) {
	i.ObserverList = append(i.ObserverList, o)
}

//从观察者列表删除
func (i *Product) Deregister(o Observer) {
	i.ObserverList = RemoveFromslice(i.ObserverList, o)
}

//通知所有用户
func (i *Product) NotifyAll() {
	for _, Observer := range i.ObserverList {
		Observer.Update(i.name)
	}
}

func RemoveFromslice(ObserverList []Observer, ObserverToRemove Observer) []Observer {
	ObserverListLength := len(ObserverList)
	for i, Observer := range ObserverList {
		if ObserverToRemove.GetID() == Observer.GetID() {
			ObserverList[ObserverListLength-1], ObserverList[i] = ObserverList[i], ObserverList[ObserverListLength-1]
			return ObserverList[:ObserverListLength-1]
		}
	}
	return ObserverList
}
