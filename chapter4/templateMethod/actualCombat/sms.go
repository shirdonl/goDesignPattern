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

package actualCombat

import "fmt"

//短信类
type Sms struct {
	Otp
}

func (s *Sms) GenRandomOTP(len int) string {
	randomOTP := "1688"
	fmt.Printf("SMS: 生成随机验证码：%s\n", randomOTP)
	return randomOTP
}

func (s *Sms) SaveOTPCache(otp string) {
	fmt.Printf("SMS: 保存验证码：%s 到缓存\n", otp)
}

func (s *Sms) GetMessage(otp string) string {
	return "登录的短信验证码是：" + otp
}

func (s *Sms) SendNotification(message string) error {
	fmt.Printf("SMS: 发送消息：%s\n", message)
	return nil
}

func (s *Sms) Publish() {
	fmt.Printf("SMS: 发布完成\n")
}
