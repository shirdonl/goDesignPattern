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

//文件类型
type File struct {
	Name string
}

//打印
func (f *File) Print(indentation string) {
	fmt.Println(indentation + f.Name)
}

//克隆
func (f *File) Clone() InterfaceNode {
	return &File{Name: f.Name + "_Clone"}
}
