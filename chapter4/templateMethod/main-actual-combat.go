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
	"github.com/shirdonl/goDesignPattern/chapter4/templateMethod/actualCombat"
)

func main() {
	//创建短信对象
	smsOTP := &actualCombat.Sms{}
	o := actualCombat.Otp{
		IOtp: smsOTP,
	}
	//生成短信验证码并发送
	o.GenAndSendOTP(4)

	fmt.Println("")
	//创建邮件对象
	EmailOTP := &actualCombat.Email{}
	o = actualCombat.Otp{
		IOtp: EmailOTP,
	}
	//生成邮件验证码并发送
	o.GenAndSendOTP(4)
}

//$ go run main-actual-combat.go
//SMS: 生成随机验证码：1688
//SMS: 保存验证码：1688 到缓存
//SMS: 发送消息：登录的短信验证码是：1688
//SMS: 发布完成
//
//EMAIL: 生成随机验证码：3699
//EMAIL: 保存验证码：3699 到缓存
//EMAIL: 发送消息：登录的短信验证码是：3699
//EMAIL:发布完成
