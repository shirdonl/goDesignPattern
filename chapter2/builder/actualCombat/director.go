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

//主管类型
type Director struct {
	Builder InterfaceBuilder
}

func NewDirector(b InterfaceBuilder) *Director {
	return &Director{
		Builder: b,
	}
}

func (d *Director) SetBuilder(b InterfaceBuilder) {
	d.Builder = b
}

func (d *Director) BuildCar() Car {
	d.Builder.SetEngineType()
	d.Builder.SetSeatsType()
	d.Builder.SetNumber()
	return d.Builder.GetCar()
}
