//++++++++++++++++++++++++++++++++++++++++
// 《Go语言设计模式》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goDesignPattern
// 仓库地址：https://github.com/shirdonl/goDesignPattern
// 交流咨询，请关注公众号"源码大数据"
//++++++++++++++++++++++++++++++++++++++++

package actualCombat

import "fmt"

//电子产品工厂
type InterfaceElectronicFactory interface {
	MakePhone() AbstractPhone
	MakeComputer() AbstractComputer
}

//获取电子产品工厂对象
func GetElectronicFactory(brand string) (InterfaceElectronicFactory, error) {
	if brand == "Xiaomi" {
		return &XiaomiFactory{}, nil
	}

	if brand == "Lenovo" {
		return &LenovoFactory{}, nil
	}

	return nil, fmt.Errorf("%s", "error brand type")
}
