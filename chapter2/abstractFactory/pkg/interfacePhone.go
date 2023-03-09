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

package pkg

//定义手机接口
type InterfacePhone interface {
	SetColor(color string)
	SetSize(size int)
	GetColor() string
	GetSize() int
}

type Phone struct {
	color string
	size  int
}

func (s *Phone) SetColor(color string) {
	s.color = color
}

func (s *Phone) GetColor() string {
	return s.color
}

func (s *Phone) SetSize(size int) {
	s.size = size
}

func (s *Phone) GetSize() int {
	return s.size
}
