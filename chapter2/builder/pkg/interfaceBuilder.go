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

//生成器接口
type InterfaceBuilder interface {
	SetSeatsType()
	SetEngineType()
	SetNumber()
	GetCar() Car
}

//获取生成器
func GetBuilder(BuilderType string) InterfaceBuilder {
	if BuilderType == "mpv" {
		return &MpvBuilder{}
	}

	if BuilderType == "suv" {
		return &SuvBuilder{}
	}
	return nil
}