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
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

const (
	minTitleLength = 1
	maxTitleLength = 50
)

var (
	// 给出无效标题时使用的错误
	ErrInvalidTitle  = errors.New("tab: could not use invalid title")
	ErrTitleTooShort = fmt.Errorf("%s: min length allowed is %d", ErrInvalidTitle, minTitleLength)
	ErrTitleTooLong  = fmt.Errorf("%s: max length allowed is %d", ErrInvalidTitle, maxTitleLength)
)

// Title 代表标签标题
type Title string

// NewTitle 返回标题和错误
func NewTitle(d string) (Title, error) {
	switch l := len(strings.TrimSpace(d)); {
	case l < minTitleLength:
		return "", ErrTitleTooShort
	case l > maxTitleLength:
		return "", ErrTitleTooLong
	default:
		return Title(d), nil
	}
}

// String 返回标题的字符串表示
func (t Title) String() string {
	return string(t)
}

// Equals 如果标题相等，则返回 true
func (t Title) Equals(t2 Title) bool {
	return t.String() == t2.String()
}

//添加标签请求
type addTabReq struct {
	Title Title `json:"tab_title"`
}

//解码JSON请求正文
func (r *addTabReq) UnmarshalJSON(data []byte) error {
	type clone addTabReq
	var req clone
	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}

	var err error
	if r.Title, err = NewTitle(req.Title.String()); err != nil {
		return err
	}

	return nil
}
