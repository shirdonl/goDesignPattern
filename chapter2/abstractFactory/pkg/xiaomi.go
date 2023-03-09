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

//小米品牌工厂
type Xiaomi struct {
}

//生产手机
func (a *Xiaomi) MakePhone() InterfacePhone {
	return &XiaomiPhone{
		Phone: Phone{
			color: "Xiaomi",
			size:  5,
		},
	}
}

//生产电脑
func (a *Xiaomi) MakeComputer() InterfaceComputer {
	return &XiaomiComputer{
		Computer: Computer{
			color: "Xiaomi",
			size:  14,
		},
	}
}
