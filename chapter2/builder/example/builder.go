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

package example

// 主管
type Director struct {
	builder Builder
}

// 初始化主管对象
func NewDirector(builder Builder) Director {
	return Director{builder}
}

// 通过一系列步骤生成产品
func (d *Director) Construct() {
	d.builder.Build()
}

// 生成器接口
type Builder interface {
	Build()
}

//具体生成器，用于构建产品的生成器
type ConcreteBuilder struct {
	result Product
}

// 初始化具体生成器对象
func NewConcreteBuilder() ConcreteBuilder {
	return ConcreteBuilder{result: Product{}}
}

// 生成产品
func (b *ConcreteBuilder) Build() {
	b.result = Product{}
}

// 返回在生成步骤中生成的产品
func (b *ConcreteBuilder) GetResult() Product {
	return Product{true}
}

// 产品
type Product struct {
	Built bool
}
