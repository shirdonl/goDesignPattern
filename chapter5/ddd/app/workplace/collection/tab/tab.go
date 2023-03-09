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
	"time"
)

// Tab 代表一个标签
type Tab struct {
	ID          int
	Title       Title
	Description string
	Icon        string
	Link        string
	Created     time.Time
	Updated     time.Time
}

// New 返回第一次创建的标签
func New(id int, title Title, description string, icon string, link string) *Tab {
	return &Tab{
		ID:          id,
		Title:       title,
		Description: description,
		Icon:        icon,
		Link:        link,
		Created:     time.Now(),
	}
}

// Update 使用新属性更新标签
func (t *Tab) Update(title Title, description string, icon string, link string) {
	t.Title = title
	t.Description = description
	t.Icon = icon
	t.Link = link
	t.Updated = time.Now()
}
