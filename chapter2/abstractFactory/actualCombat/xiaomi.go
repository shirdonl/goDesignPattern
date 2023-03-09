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

//小米品牌工厂
type XiaomiFactory struct {
}

//生产手机
func (a *XiaomiFactory) MakePhone() AbstractPhone {
	return &XiaomiPhone{
		Phone: Phone{
			color: "White",
			size:  5,
		},
	}
}

//生产电脑
func (a *XiaomiFactory) MakeComputer() AbstractComputer {
	return &XiaomiComputer{
		Computer: Computer{
			color: "Black",
			size:  14,
		},
	}
}
