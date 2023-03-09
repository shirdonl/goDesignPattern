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

import "fmt"

// Handler 定义了一个处理程序来处理给定的 handleID
type Handler interface {
	SetNext(handler Handler)
	Handle(handleID int) int
}

//基础处理者
type BaseHandler struct {
	name     string
	next     Handler
	handleID int
}

//NewHandler 返回一个新的处理程序
func NewBaseHandler(name string, next Handler, handleID int) Handler {
	return &BaseHandler{name, next, handleID}
}

// Handle 处理给定的 handleID
func (h *BaseHandler) Handle(handleID int) int {
	if handleID < 4 {
		ch := &ConcreteHandler{}
		ch.Handle(handleID)
		fmt.Println(h.name)

		if h.next != nil {
			h.next.Handle(handleID + 1)
		}

		return handleID + 1
	}
	return 0
}

// 设置下一个处理者
func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

//具体处理者
type ConcreteHandler struct {
}

//具体处理者的处理方法
func (ch *ConcreteHandler) Handle(handleID int) {
	fmt.Println("ConcreteHandler handleID:", handleID)
}
