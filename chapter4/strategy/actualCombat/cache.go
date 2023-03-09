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

package actualCombat

type Cache struct {
	storage       map[string]string
	AlgorithmType AlgorithmType
	capacity      int
	maxCapacity   int
}

func InitCache(e AlgorithmType) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:       storage,
		AlgorithmType: e,
		capacity:      0,
		maxCapacity:   2,
	}
}

func (c *Cache) SetAlgorithmType(e AlgorithmType) {
	c.AlgorithmType = e
}

func (c *Cache) Add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.Delete()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) Get(key string) {
	delete(c.storage, key)
}

func (c *Cache) Delete() {
	c.AlgorithmType.Delete(c)
	c.capacity--
}
