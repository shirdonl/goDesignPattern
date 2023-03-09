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

import "fmt"

//适配者类
type ObjectAdaptee struct {
}

// 目标接口
type ObjectTarget interface {
	Execute()
}

//适配者类的方法
func (b *ObjectAdaptee) SpecificExecute() {
	fmt.Println("最终执行的方法")
}

//适配器类
type ObjectAdapter struct {
	Adaptee ObjectAdaptee
}

// 适配器类的方法
func (p *ObjectAdapter) Execute() {
	p.Adaptee.SpecificExecute()
}
