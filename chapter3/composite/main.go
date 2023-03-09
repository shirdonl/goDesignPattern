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

import "github.com/shirdonl/goDesignPattern/chapter3/composite/example"

func main() {
	composite := example.NewComposite()
	leaf1 := example.NewLeaf(99)
	composite.Add(leaf1)
	leaf2 := example.NewLeaf(100)
	composite.Add(leaf2)
	leaf3 := example.NewComposite()
	composite.Add(leaf3)
	composite.Execute()
}

//$ go run main.go
//99  100
