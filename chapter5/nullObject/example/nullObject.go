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
	"errors"
	"fmt"
	"github.com/stretchr/testify/mock"
	"os"
)

//抽象对象
type AbstractObject struct {
	mock.Mock
}

//方法
func (m *AbstractObject) Request(str string) (string, error) {
	args := m.Called(str)
	return args.String(0), args.Error(1)
}

//空对象
type NullObject struct {
	AbstractObject
}

func (w *NullObject) Request(str string) (string, error) {
	return "null", errors.New("not implemented yet!")
}

//真实对象
type RealObject struct {
	AbstractObject
}

func (w *RealObject) Request(str string) (string, error) {
	return str, nil
}

func main() {
	os.Setenv("RealObject", "true")
	var realObject *RealObject
	var nullObject *NullObject
	_, realObjectFlagExists := os.LookupEnv("RealObject")
	if realObjectFlagExists {
		realObject = &RealObject{}
		fmt.Println(realObject.Request("real"))
	} else {
		nullObject = &NullObject{}
		fmt.Println(nullObject.Request("null"))
	}
}
