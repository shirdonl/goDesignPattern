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

//蓝队服装
type BlueTeamDress struct {
	color string
}

func (t *BlueTeamDress) GetColor() string {
	return t.color
}

func newBlueTeamDress() *BlueTeamDress {
	return &BlueTeamDress{color: "blue"}
}
