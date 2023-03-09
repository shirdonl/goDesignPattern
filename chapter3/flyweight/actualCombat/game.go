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

//创建游戏
type NewGame struct {
}

//创建蓝队队员
func (ng *NewGame) AddBlueTeam(DressType string) *Player {
	return NewPlayer("terrorist", DressType)
}

//创建红队队员
func (ng *NewGame) AddRedTeam(DressType string) *Player {
	return NewPlayer("counterBlueTeam", DressType)
}
