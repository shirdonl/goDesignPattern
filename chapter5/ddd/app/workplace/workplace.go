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

package workspace

import (
	"github.com/shirdonl/goDesignPattern/chapter6/ddd/app/workplace/collection"
	"time"
)

// 工作区
type Workspace struct {
	ID          int
	Name        string
	CustomerID  int
	Collections []*collection.Collection
	Created     time.Time
	Updated     time.Time
}

// 返回第一次创建的工作区
func New(id int, name string, customerID int) *Workspace {
	return &Workspace{
		ID:          id,
		Name:        name,
		CustomerID:  customerID,
		Collections: make([]*collection.Collection, 0),
		Created:     time.Now(),
	}
}

// 更改工作区的名称
func (w *Workspace) Rename(name string) {
	w.Name = name
	w.Updated = time.Now()
}

// 添加一个集合
func (w *Workspace) AddCollections(collections ...*collection.Collection) {
	w.Collections = append(w.Collections, collections...)
	w.Updated = time.Now()
}

// 如果集合存在，则删除它
func (w *Workspace) RemoveCollection(id int) bool {
	for i, coll := range w.Collections {
		if coll.ID == id {
			w.Collections[i] = w.Collections[len(w.Collections)-1]
			w.Collections[len(w.Collections)-1] = nil
			w.Collections = w.Collections[:len(w.Collections)-1]
			w.Updated = time.Now()
			return true
		}
	}

	return false
}

// 重命名集合（如果存在）
func (w *Workspace) RenameCollection(id int, name string) bool {
	for _, coll := range w.Collections {
		if coll.ID == id {
			coll.Rename(name)
			w.Updated = time.Now()
			return true
		}
	}

	return false
}
