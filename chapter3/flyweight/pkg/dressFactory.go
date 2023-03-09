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

const (
	//蓝队服装类型
	BlueTeamDressType = "Blue Dress"
	//红队服装类型
	RedTeamDressType = "Red Dress"
)

var (
	DressFactorySingleInstance = &DressFactory{
		DressMap: make(map[string]Dress),
	}
)

//享元服装工厂
type DressFactory struct {
	DressMap map[string]Dress
}

//获取服装类型
func (d *DressFactory) GetDressByType(DressType string) (Dress, error) {
	if d.DressMap[DressType] != nil {
		return d.DressMap[DressType], nil
	}

	if DressType == BlueTeamDressType {
		d.DressMap[DressType] = newBlueTeamDress()
		return d.DressMap[DressType], nil
	}
	if DressType == RedTeamDressType {
		d.DressMap[DressType] = newRedTeamDress()
		return d.DressMap[DressType], nil
	}

	return nil, fmt.Errorf("%s","Wrong Dress type")
}

//获取服装工厂单例
func GetDressFactorySingleInstance() *DressFactory {
	return DressFactorySingleInstance
}
