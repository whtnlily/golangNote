package tool

import (
	"log"
	"sync"
)

//懒汉式——非线程安全 单例模式
//使用结构体代替类
type Tool struct {
	values int

}
//建立私有变量
var instance *Tool
//锁对象
var lock sync.Mutex
//获取单例对象的方法，引用传递返回
func GetInstance() *Tool {
	//加锁保证线程安全,在非线程安全的基本上，利用 Sync.Mutex 进行加锁保证线程安全，但由于每次调用该方法都进行了加锁操作，在性能上不是很高效。
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = new(Tool)
	}
	return instance
}

func (t *Tool) Test()  {
	log.Println(">>> 懒汉式单例模式")
}

