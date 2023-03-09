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

//定义需要被适配的类
type ObjectAdaptee struct {
}

// Target 是要适配的目标接口
type ObjectTarget interface {
	Execute()
}

//定义了用于执行的方法SpecificExecute()
func (b *ObjectAdaptee) SpecificExecute() {
	fmt.Println("最终执行的方法")
}

//Adapter 是新接口 Target 的适配器，通过关联 Adaptee 类来实现
type ObjectAdapter struct {
	Adaptee ObjectAdaptee
}

// 同时继承了 Adaptee 类
func (p *ObjectAdapter) Execute() {
	p.Adaptee.SpecificExecute()
}
