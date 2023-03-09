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

//SUV生成器
type SuvBuilder struct {
	SeatsType string
	EngineType   string
	Number      int
}

func newSuvBuilder() *SuvBuilder {
	return &SuvBuilder{}
}

func (b *SuvBuilder) SetSeatsType() {
	b.SeatsType = "SUV型座椅"
}

func (b *SuvBuilder) SetEngineType() {
	b.EngineType = "SUV型引擎"
}

func (b *SuvBuilder) SetNumber() {
	b.Number = 6
}

func (b *SuvBuilder) GetCar() Car {
	return Car{
		EngineType:   b.EngineType,
		SeatsType: b.SeatsType,
		Number:      b.Number,
	}
}
