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

package pkg

//定义真实主体类
type Application struct {
}

//处理请求
func (a *Application) HandleRequest(url, method string) (int, string) {
	if url == "/user/status" && method == "GET" {
		return 200, "Ok"
	}

	if url == "/user/login" && method == "POST" {
		return 201, "User Login"
	}
	return 404, "Not Ok"
}
