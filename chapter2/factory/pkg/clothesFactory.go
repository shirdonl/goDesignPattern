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

func MakeClothes(clothesType string) (IClothes, error) {
	if clothesType == "ANTA" {
		return newANTA(), nil
	}
	if clothesType == "PEAK" {
		return newPEAK(), nil
	}
	return nil, fmt.Errorf("Wrong clothes type passed")
}
