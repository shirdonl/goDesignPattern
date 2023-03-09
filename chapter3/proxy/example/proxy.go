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

package example

import (
	"fmt"
)

// 服务接口
type ServiceInterface interface {
	Execute(access string)
}

// 服务实现了用于执行任务的 ServiceInterface 接口
type Service struct {
}

// 服务对象的方法
func (t *Service) Execute(access string) {
	fmt.Println("Proxy Service: " + access)
}

// 代理对象
type Proxy struct {
	realService *Service
}

// 创建代理对象
func NewProxy() *Proxy {
	return &Proxy{realService: &Service{}}
}

// 拦截 Execute 命令并将其重新路由到服务命令
func (t *Proxy) Execute(access string) {
	if access == "yes" {
		t.realService.Execute(access)
	}
}
