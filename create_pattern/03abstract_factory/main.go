package main

import "fmt"

// 接口
type AbstrarctApple interface {
	ShowApple()
}
type AbstrarctBanana interface {
	ShowBanana()
}
type AbstrarctPear interface {
	ShowPear()
}

type AbstractFactory interface {
	CreateApple() AbstrarctApple
	CreateBanana() AbstrarctBanana
	CreatePear() AbstrarctPear
}

// 实现
type ChinaApple struct{}

func (ca *ChinaApple) ShowApple() {
	fmt.Println("ChinaApple")
}

type ChinaBanana struct{}

func (cb *ChinaBanana) ShowBanana() {
	fmt.Println("ChinaBanana")
}

type ChinaPear struct{}

func (cp *ChinaPear) ShowPear() {
	fmt.Println("ChinaPear")

}

type ChineseFatory struct{}

func (factory *ChineseFatory) CreateApple() AbstrarctApple {
	var apple AbstrarctApple
	apple = new(ChinaApple)
	return apple
}
func (factory *ChineseFatory) CreateBanana() AbstrarctBanana {
	var banana AbstrarctBanana
	banana = new(ChinaBanana)
	return banana
}
func (factory *ChineseFatory) CreatePear() AbstrarctPear {
	var pear AbstrarctPear
	pear = new(ChinaPear)
	return pear
}
