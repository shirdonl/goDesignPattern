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

//装饰
type Decorator struct {
	Phone Phone
}

//装饰设置组件方法
func (d *Decorator) SetComponent(c Phone) {
	d.Phone = c
}

//装饰方法
func (d *Decorator) GetPrice() {
	if d.Phone != nil {
		d.Phone.GetPrice()
	}
}
