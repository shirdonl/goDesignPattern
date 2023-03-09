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

//定义用户模型
type User struct {
	Id string
}

//更新
func (c *User) Update(ProductName string) {
	fmt.Printf("发送邮件给用户： %s ，商品： %s 上架啦～ \n", c.Id, ProductName)
}

//获取编号
func (c *User) GetID() string {
	return c.Id
}
