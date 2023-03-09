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
	"github.com/shirdonl/goDesignPattern/chapter3/proxy/actualCombat"
)

func main() {
	//初始化Apache服务器
	ApacheServer := actualCombat.NewApacheServer()
	userStatusURL := "/user/status"
	userLoginURL := "/user/login"

	//发送一个GET请求
	httpCode, body := ApacheServer.HandleRequest(userStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", userStatusURL, httpCode, body)

	//发送一个POST请求
	httpCode, body = ApacheServer.HandleRequest(userStatusURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", userStatusURL, httpCode, body)

	//发送一个GET请求
	httpCode, body = ApacheServer.HandleRequest(userLoginURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", userStatusURL, httpCode, body)

	//发送一个POST请求
	httpCode, body = ApacheServer.HandleRequest(userLoginURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", userStatusURL, httpCode, body)
}

//$ go run main-actual-combat.go
//
//Url: /user/status
//HttpCode: 200
//Body: Ok
//
//Url: /user/status
//HttpCode: 404
//Body: Not Ok
//
//Url: /user/status
//HttpCode: 201
//Body: User Login
//
//Url: /user/status
//HttpCode: 404
//Body: Not Ok
