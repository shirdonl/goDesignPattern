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
	"github.com/shirdonl/goDesignPattern/chapter2/objectPool/example"
)

func main() {
	num := func() interface{} {
		return 10.0
	}
	pool := example.NewPool(num)
	object := pool.Acquire()

	fmt.Println(pool.Inuse)
	fmt.Println(pool.Available)

	pool.Release(object)
	fmt.Println(pool.Inuse)
	fmt.Println(pool.Available)
}

//$ go run main.go
//[10]
//[]
//[]
//[10]
