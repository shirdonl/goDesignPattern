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

package tab

import (
	"fmt"
	"github.com/shirdonl/goDesignPattern/chapter5/ddd/app/workplace/collection/tab"
)

//这里是一个 MySQL 实现
type MysqlRepo struct {
}

func (r *MysqlRepo) Add(t *tab.Tab) error {
	// ...
	return fmt.Errorf("error: %s %s", tab.ErrRepoAdd, "a more detailed reason here")
}
