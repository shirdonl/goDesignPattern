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

import "errors"

var (
	//存储库返回的错误
	ErrRepoNextID = errors.New("tab: could not return next id")
	ErrRepoList   = errors.New("tab: could not list")
	ErrNotFound   = errors.New("tab: could not find")
	ErrRepoGet    = errors.New("tab: could not get")
	ErrRepoAdd    = errors.New("tab: could not add")
	ErrRepoRemove = errors.New("tab: could not remove")
)

type ID int

type Repo interface {
	//NextID 返回下一个空闲 ID 并在失败时返回错误
	NextID() (ID, error)
	// 列表返回一个选项卡切片并在失败的情况下返回一个错误
	List() ([]*Tab, error)
	// Find 返回一个 tab 或者 nil ，如果它没有找到并且在失败的情况下，则返回一个错误
	Find(ID) (*Tab, error)
	// 如果未找到或失败，则获取返回选项卡和错误
	Get(ID) (*Tab, error)
	// 添加持久化选项卡（已经存在或不存在）并在失败时返回错误
	Add(*Tab) error
	// 删除选项卡并返回并在未找到或失败的情况下出错
	Remove(ID) error
}
