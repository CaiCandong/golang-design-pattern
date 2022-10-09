package main

import "fmt"

// 饿汉模式
type hunger_singelton struct{}

func (s *hunger_singelton) DoSomeThing() {
	fmt.Println("单例对象的方法")
}

var hunger_instance *hunger_singelton = new(hunger_singelton)

func GetHungerInstance() *hunger_singelton {
	return hunger_instance
}

// func main() {
// 	s := GetInstance()
// 	s.DoSomeThing()
// }
