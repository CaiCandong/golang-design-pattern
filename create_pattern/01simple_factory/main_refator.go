package main

import "fmt"

type Fruit interface {
	Show()
}
type Apple struct {
}

func (apple *Apple) Show() {
	fmt.Println("我是苹果")
}

type Banana struct {
}

func (banana *Banana) Show() {
	fmt.Println("我是香蕉")
}

type Pear struct {
	Fruit
}

func (pear *Pear) Show() {
	fmt.Println("我是梨")
}

type Factory struct{}

//  违反“开闭原则”，添加新产品需要修改工厂逻辑，工厂越来越复杂。
//  增加新产品时，该方法的源代码需要进行修改
//  开闭原则：类的改动是通过增加代码进行的，而不是修改源代码。
func (fac *Factory) CreateFruit(kind string) Fruit {
	var fruit Fruit

	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = new(Banana)
	} else if kind == "pear" {
		fruit = new(Pear)
	}

	return fruit
}

func main() {
	factory := new(Factory)

	apple := factory.CreateFruit("apple")
	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()
}
