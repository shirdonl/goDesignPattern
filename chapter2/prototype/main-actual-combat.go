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
	"fmt"
	"github.com/shirdonl/goDesignPattern/chapter2/prototype/actualCombat"
)

func main() {
	//声明文件对象File1
	File1 := &actualCombat.File{Name: "File1"}
	//声明文件对象File2
	File2 := &actualCombat.File{Name: "File2"}
	//声明文件对象File3
	File3 := &actualCombat.File{Name: "File3"}

	//声明文件夹对象Folder1
	Folder1 := &actualCombat.Folder{
		Children: []actualCombat.InterfaceNode{File1},
		Name:     "文件夹Folder1",
	}

	//声明文件夹对象Folder2
	Folder2 := &actualCombat.Folder{
		Children: []actualCombat.InterfaceNode{Folder1, File2, File3},
		Name:     "文件夹Folder2",
	}
	fmt.Println("\n打印文件夹Folder2的层级:")
	Folder2.Print("  ")

	CloneFolder := Folder2.Clone()
	fmt.Println("\n打印复制文件夹Folder2的层级:")
	CloneFolder.Print("  ")
}

//$ go run main-actual-combat.go
//
//打印文件夹Folder2的层级:
//  文件夹Folder2
//    文件夹Folder1
//        File1
//    File2
//    File3
//
//打印复制文件夹Folder2的层级:
//  文件夹Folder2_Clone
//    文件夹Folder1_Clone
//        File1_Clone
//    File2_Clone
//    File3_Clone
