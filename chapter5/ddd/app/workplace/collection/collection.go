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

package collection

import (
	"github.com/shirdonl/goDesignPattern/chapter5/ddd/app/workplace/collection/tab"
	"time"
)

// Collection 代表一个集合
type Collection struct {
	ID      int
	Name    string
	Tabs    []*tab.Tab
	Created time.Time
	Updated time.Time
}

// New 返回第一次创建的集合
func New(id int, name string) *Collection {
	return &Collection{
		ID:      id,
		Name:    name,
		Tabs:    make([]*tab.Tab, 0),
		Created: time.Now(),
	}
}

// Rename重命名集合
func (c *Collection) Rename(name string) {
	c.Name = name
	c.Updated = time.Now()
}

// AddTabs 将Tab添加到集合
func (c *Collection) AddTabs(tabs ...*tab.Tab) {
	c.Tabs = append(c.Tabs, tabs...)
	c.Updated = time.Now()
}

// RemoveTab 如果标签存在，则删除它
func (c *Collection) RemoveTab(id int) bool {
	for i, t := range c.Tabs {
		if t.ID == id {
			c.Tabs[i] = c.Tabs[len(c.Tabs)-1]
			c.Tabs[len(c.Tabs)-1] = nil
			c.Tabs = c.Tabs[:len(c.Tabs)-1]
			c.Updated = time.Now()
			return true
		}
	}

	return false
}

// FindTab 如果存在则返回一个标签
func (c *Collection) FindTab(id int) (*tab.Tab, bool) {
	for _, t := range c.Tabs {
		if t.ID == id {
			return t, true
		}
	}

	return nil, false
}

// UpdateTab 更新标签（如果存在）
func (c *Collection) UpdateTab(t *tab.Tab) bool {
	for i, tb := range c.Tabs {
		if tb.ID == t.ID {
			c.Tabs[i] = t
			c.Updated = time.Now()
			return true
		}
	}

	return false
}
