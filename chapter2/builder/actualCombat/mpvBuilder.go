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

//MPV生成器
type MpvBuilder struct {
	SeatsType  string
	EngineType string
	Number     int
}

func NewMpvBuilder() *MpvBuilder {
	return &MpvBuilder{}
}

func (b *MpvBuilder) SetSeatsType() {
	b.SeatsType = "MPV型座椅"
}

func (b *MpvBuilder) SetEngineType() {
	b.EngineType = "MPV型引擎"
}

func (b *MpvBuilder) SetNumber() {
	b.Number = 8
}

func (b *MpvBuilder) GetCar() Car {
	return Car{
		EngineType: b.EngineType,
		SeatsType:  b.SeatsType,
		Number:     b.Number,
	}
}
