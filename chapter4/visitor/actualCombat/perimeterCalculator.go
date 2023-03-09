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

import (
	"fmt"
)

//具体访问者——周长计算器
type PerimeterCalculator struct {
	perimeter float32
}

func (a *PerimeterCalculator) VisitForSquare(s *Square) {
	var perimeter float32
	perimeter = 4 * s.Side
	fmt.Printf("计算正方形的周长：%f \n", perimeter)
}

func (a *PerimeterCalculator) VisitForCircle(c *Circle) {
	var perimeter float32
	perimeter = 3.14 * 2 * c.Radius
	fmt.Printf("计算圆形的周长：%f \n", perimeter)
}

func (a *PerimeterCalculator) VisitForRectangle(r *Rectangle) {
	var perimeter float32
	perimeter = 2*r.B + 2*r.L
	fmt.Printf("计算矩形的周长：%f \n", perimeter)
}
