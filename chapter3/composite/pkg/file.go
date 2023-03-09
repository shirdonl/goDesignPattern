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

import "fmt"

type File struct {
	Name string
}

func (f *File) Search(keyword string) {
	fmt.Printf("在文件 %s 中递归搜索关键 %s \n", f.Name, keyword)
}

func (f *File) GetName() string {
	return f.Name
}