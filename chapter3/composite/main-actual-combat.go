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

package main

import (
	"github.com/shirdonl/goDesignPattern/chapter3/composite/actualCombat"
)

func main() {
	File1 := &actualCombat.File{Name: "File1"}
	File2 := &actualCombat.File{Name: "File2"}
	File3 := &actualCombat.File{Name: "File3"}

	Folder1 := &actualCombat.Folder{
		Name: "Folder1",
	}

	Folder1.Add(File1)

	Folder2 := &actualCombat.Folder{
		Name: "Folder2",
	}
	Folder2.Add(File2)
	Folder2.Add(File3)
	Folder2.Add(Folder1)

	Folder2.Search("keyword")
}

//$ go run main-actual-combat.go
//在文件夹 Folder2 中递归搜索关键 keyword
//在文件 File2 中递归搜索关键 keyword
//在文件 File3 中递归搜索关键 keyword
//在文件夹 Folder1 中递归搜索关键 keyword
//在文件 File1 中递归搜索关键 keyword
