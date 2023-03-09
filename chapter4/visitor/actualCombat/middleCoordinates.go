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

//中点坐标计算器
type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) VisitForSquare(s *Square) {
	//...省略具体逻辑
	fmt.Println("计算正方形的中点坐标")
}

func (a *MiddleCoordinates) VisitForCircle(c *Circle) {
	//...省略具体逻辑
	fmt.Println("计算圆形的中点坐标")
}

func (a *MiddleCoordinates) VisitForRectangle(t *Rectangle) {
	//...省略具体逻辑
	fmt.Println("计算矩形的中点坐标")
}
