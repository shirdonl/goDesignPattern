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

import (
	"fmt"
)

// 定义被适配的类
type Adaptee struct {
}

// 目标接口
type Target interface {
	Execute()
}

//定义了用于执行的方法SpecificExecute()
func (a *Adaptee) SpecificExecute() {
	fmt.Println("最终执行的方法")
}

// Adapter 是新接口 Target 的适配器，继承了 Adaptee 类
type Adapter struct {
	*Adaptee
}

// 实现 Target 接口，同时继承了 Adaptee 类
func (a *Adapter) Execute() {
	a.SpecificExecute()
}
