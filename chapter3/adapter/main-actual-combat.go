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

import "github.com/shirdonl/goDesignPattern/chapter3/adapter/actualCombat"

func main() {
	//创建客户端
	Client := &actualCombat.Client{}
	Mac := &actualCombat.Mac{}

	//客户端插入Lightning类型连接器到Mac电脑
	Client.InsertIntoComputer(Mac)

	WindowsAdapter := &actualCombat.Windows{}
	WindowsAdapterAdapter := &actualCombat.Adapter{
		WindowsMachine: WindowsAdapter,
	}

	//客户端插入Lightning类型连接器到Windows适配器
	Client.InsertIntoComputer(WindowsAdapterAdapter)
}

//$ go run main-actual-combat.go
//Lightning类型接口已插入Mac电脑
//客户端将Lightning类型接口插入计算机
//适配器将Lightning类型信号转换为USB
//USB接口已插入Windows电脑
