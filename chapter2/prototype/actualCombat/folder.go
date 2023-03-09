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

import "fmt"

//文件夹
type Folder struct {
	Children []InterfaceNode
	Name     string
}

//打印
func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.Name)
	for _, i := range f.Children {
		i.Print(indentation + indentation)
	}
}

//克隆
func (f *Folder) Clone() InterfaceNode {
	CloneFolder := &Folder{Name: f.Name + "_Clone"}
	var tempChildren []InterfaceNode
	for _, i := range f.Children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)
	}
	CloneFolder.Children = tempChildren
	return CloneFolder
}
