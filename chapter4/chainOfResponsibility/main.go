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
	"github.com/shirdonl/goDesignPattern/chapter4/chainOfResponsibility/example"
)

func main() {
	barry := example.NewBaseHandler("Barry", nil, 1)
	shirdon := example.NewBaseHandler("Shirdon", barry, 2)
	jack := example.NewBaseHandler("Shirdon", shirdon, 3)
	res := shirdon.Handle(2)
	res1 := jack.Handle(3)
	fmt.Println(res)
	fmt.Println(res1)
}

//$ go run main.go
//ConcreteHandler handleID: 2
//Shirdon
//ConcreteHandler handleID: 3
//Barry
//ConcreteHandler handleID: 3
//Shirdon
//3
//4
