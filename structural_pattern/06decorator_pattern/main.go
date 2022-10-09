package main

import "fmt"

type Phone interface {
	Show()
}

// 手机贴膜装饰器
type MoDecorator struct {
	Phone
}

func (d *MoDecorator) Show() {
	d.Phone.Show()
	fmt.Println("手机贴膜")
}
func NewMoDecorator(p Phone) Phone {
	return &MoDecorator{p}
}

// 手机壳装饰器
type KeDecorator struct {
	Phone
}

func (d *KeDecorator) Show() {
	d.Phone.Show()
	fmt.Println("装一个手机壳")
}
func NewKeDecorator(p Phone) Phone {
	return &KeDecorator{p}
}

// 基础手机
type Huawei struct{}

func (d *Huawei) Show() {
	fmt.Println("亮出我的华为手机")
}

func main() {
	var huawei Phone
	huawei = &Huawei{}
	huawei = NewKeDecorator(NewMoDecorator(huawei))
	huawei.Show()
}
