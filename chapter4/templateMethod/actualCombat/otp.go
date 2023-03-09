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

//定义一次性密码接口
type IOtp interface {
	GenRandomOTP(int) string
	SaveOTPCache(string)
	GetMessage(string) string
	SendNotification(string) error
	Publish()
}

//定义一次性密码类
type Otp struct {
	IOtp IOtp
}

//生成验证码并发送
func (o *Otp) GenAndSendOTP(otpLength int) error {
	//生成随机验证码
	otp := o.IOtp.GenRandomOTP(otpLength)
	o.IOtp.SaveOTPCache(otp)
	message := o.IOtp.GetMessage(otp)
	err := o.IOtp.SendNotification(message)
	if err != nil {
		return err
	}
	o.IOtp.Publish()
	return nil
}
