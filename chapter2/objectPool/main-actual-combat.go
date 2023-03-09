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
	"github.com/shirdonl/goDesignPattern/chapter2/objectPool/actualCombat"
	"log"
	"sync"
)

func main() {
	// 初始化一个包含五个资源的资源池
	// 可以调整为 1 或 10 以查看差异
	size := 5
	p := actualCombat.New(size)

	// 调用资源池
	doWork := func(workId int, wg *sync.WaitGroup) {
		defer wg.Done()
		// 从资源池中获取资源对象
		res, err := p.GetResource()
		if err != nil {
			log.Println(err)
			return
		}
		//返回的资源对象
		defer p.GiveBackResource(res)
		// 使用资源处理工作
		res.Do(workId)
	}

	// 模拟100个并发进程从资产池中获取资源对象
	num := 100
	wg := new(sync.WaitGroup)
	wg.Add(num)
	for i := 0; i < num; i++ {
		go doWork(i, wg)
	}
	wg.Wait()
}

//$ go run main-actual-combat.go
//2022/06/04 16:38:39 using resource #1 finished work 3 finish
//2022/06/04 16:38:39 using resource #0 finished work 5 finish
//2022/06/04 16:38:39 using resource #0 finished work 8 finish
//2022/06/04 16:38:39 using resource #0 finished work 12 finish
//2022/06/04 16:38:39 using resource #3 finished work 9 finish
//2022/06/04 16:38:39 using resource #3 finished work 11 finish
//2022/06/04 16:38:39 using resource #2 finished work 1 finish
//...省略更多记录
