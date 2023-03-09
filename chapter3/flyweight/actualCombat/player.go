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

//队员类
type Player struct {
	Dress      Dress
	PlayerType string
	lat        int
	long       int
}

//创建一个队员
func NewPlayer(PlayerType, DressType string) *Player {
	Dress, _ := GetDressFactorySingleInstance().GetDressByType(DressType)
	return &Player{
		PlayerType: PlayerType,
		Dress:      Dress,
	}
}

//创建队员位置
func (p *Player) NewLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
