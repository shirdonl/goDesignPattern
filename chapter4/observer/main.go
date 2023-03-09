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

package main

import "github.com/shirdonl/goDesignPattern/chapter4/observer/example"

func main() {
	//event := example.Event{"event"}
	//observer := example.NewObserver("Barry")
	//observer.Update(event)

	notifier := example.Subject{}
	observers := []example.Observer{
		example.NewObserver("Barry"),
		example.NewObserver("Jack"),
		example.NewObserver("Shirdon"),
	}

	for i := 0; i < len(observers); i++ {
		notifier.Register(observers[i])
	}
	notifier.Unregister(observers[1])
	notifier.NotifyObservers(example.Event{"Received an email!"})
}

//$ go run main.go
//ConcreteObserver 'Barry' received event 'Received an email!'
//ConcreteObserver 'Shirdon' received event 'Received an email!'
