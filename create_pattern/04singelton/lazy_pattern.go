package main

import (
	"fmt"
	"sync"
)

type singelton struct{}

func (s *singelton) SomeThing() {
	fmt.Println("单例对象的某方法")
}

var instance *singelton

// 1.并发不安全
/*
func GetInstance() *singelton {
	if instance == nil {
		instance= new(singelton)
	}
	return instance
}
*/

// 2.使用锁保证并发安全（损失性能)
// 定义锁
/*
var lock sync.Mutex

func GetInstance() *singelton {
	//为了线程安全，增加互斥
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance= new(singelton)
	}
	return instance
}
*/
// 3.使用原子操作提升性能
/*
var mark uint32
var lock sync.Mutex

func GetInstance() *singelton {
	if atomic.LoadUint32(&mark) == 1 {
		return instance
	}
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		atomic.StoreUint32(&mark, 1)
		instance = new(singelton)
	}
	return instance
}
*/
// 4.使用Once进行优化
// /*
var once sync.Once

func GetInstance() *singelton {
	once.Do(func() {
		instance = new(singelton)
	})
	return instance
}

// */
func main() {
	s := GetInstance()
	s.SomeThing()
}
