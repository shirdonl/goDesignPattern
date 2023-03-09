//++++++++++++++++++++++++++++++++++++++++
// 《Go语言设计模式》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goDesignPattern
// 仓库地址：https://github.com/shirdonl/goDesignPattern
// 交流咨询，请关注公众号"源码大数据"
//++++++++++++++++++++++++++++++++++++++++

package actualCombat

//定义电脑接口
type AbstractComputer interface {
	SetColor(color string)
	SetSize(size int)
	GetColor() string
	GetSize() int
}

type Computer struct {
	color string
	size  int
}

func (s *Computer) SetColor(color string) {
	s.color = color
}

func (s *Computer) GetColor() string {
	return s.color
}

func (s *Computer) SetSize(size int) {
	s.size = size
}

func (s *Computer) GetSize() int {
	return s.size
}
