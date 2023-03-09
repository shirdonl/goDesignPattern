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

import "fmt"

type singleton struct {
}

var instance *singleton

func init() {
	if instance == nil {
		instance = new(singleton)
		fmt.Println("创建单个实例")
	}
}

// 提供获取实例的方法
func GetInstance() *singleton {
	return instance
}

func main() {
	for i := 0; i < 3; i++ {
		GetInstance()
	}
}
