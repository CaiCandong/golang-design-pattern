package main

import "fmt"

type V5 interface {
	Use5V()
}

type Phone struct {
	v V5 //依赖5v充电器
}

func (p *Phone) Charge() {
	fmt.Println("Phone进行充电...")
	p.v.Use5V()
}
func NewPhone(v V5) *Phone {
	return &Phone{v}
}

type V220 struct{}

func (v *V220) Use220V() {
	fmt.Println("使用220v的电压")
}

//电源适配器
type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	fmt.Println("使用适配器进行充电")

	//调用适配者的方法
	a.v220.Use220V()
}

func NewAdapter(v220 *V220) V5 {
	var adapter V5 = &Adapter{v220}
	return adapter
}

// ------- 业务逻辑层 -------
func main() {
	iphone := NewPhone(NewAdapter(new(V220)))

	iphone.Charge()
}
