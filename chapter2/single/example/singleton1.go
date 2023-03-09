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

import (
	"fmt"
	"sync"
)

//单例类
type singleton struct {
	value int
}

//声明私有变量
var instance *singleton

//
//
////声明私有变量
//var instance *singleton
////获取单例对象
//func GetInstance() *singleton {
//	if instance == nil {
//		instance = new(singleton)
//	}
//	return instance
//}
//
////声明私有变量
//var instance *singleton
//
////声明锁对象
//var mutex sync.Mutex
////加锁保证协程安全
//func GetInstance() *singleton {
//	mutex.Lock()
//	defer mutex.Unlock()
//	if instance == nil {
//		instance = new(singleton)
//	}
//	return instance
//}

////声明锁对象
//var mutex sync.Mutex
////当对象为空时，对对象加锁，当创建好对象后，获取对象时就不用加锁了
//func GetInstance() *singleton {
//	if instance == nil {
//		mutex.Lock()
//		if instance == nil {
//			instance = new(singleton)
//			fmt.Println("创建单个实例")
//		}
//		mutex.Unlock()
//	}
//	return instance
//}

var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = new(singleton)
		fmt.Println("创建单个实例")
	})
	return instance
}

func main() {
	for i := 0; i < 3; i++ {
		GetInstance()
	}
}
