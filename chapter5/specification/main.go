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
	"github.com/shirdonl/goDesignPattern/chapter5/specification/example"
)

func main() {
	//声明业务规格对象1
	biz1 := example.NewBusinessSpecification()
	//声明业务规格对象2
	biz2 := example.NewBusinessSpecification()

	andResult := biz1.And(biz2)

	object := example.Object{
		Attribute: 8,
	}

	// 检查规格
	result := andResult.IsSatisfiedBy(object)
	fmt.Println(result)
}

//$ go run main.go
//true
