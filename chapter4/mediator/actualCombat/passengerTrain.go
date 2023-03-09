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

//客运火车类
type PassengerTrain struct {
	Mediator Mediator
}

//火车到达
func (p *PassengerTrain) Arrive() {
	if !p.Mediator.CanArrive(p) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

//火车离开
func (p *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain: Leaving")
	p.Mediator.NotifyAboutDeparture()
}

//允许到达
func (p *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	p.Arrive()
}
