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

//验证码
type VerificationCode struct {
	code int
}

//创建验证码
func NewVerificationCode(code int) *VerificationCode {
	return &VerificationCode{
		code: code,
	}
}

//检查验证码
func (s *VerificationCode) CheckCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("%s", "验证码不正确")
	}
	fmt.Println("验证通过～")
	return nil
}
