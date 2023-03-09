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

//Windows系统
type Windows struct {
	Printer Printer
}

//打印
func (w *Windows) Print() {
	fmt.Println("Print request for Windows")
	w.Printer.PrintFile()
}

//设置打印机
func (w *Windows) SetPrinter(p Printer) {
	w.Printer = p
}
