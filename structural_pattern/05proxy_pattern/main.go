package main

import "fmt"

type Goods struct {
	Kind string
	Fact bool
}
type Shopping interface {
	Buy(goods *Goods)
}

type KoreaShopping struct{}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国进行了购物, 买了 ", goods.Kind)
}

type OverseasProxy struct {
	shopping Shopping //对原有方法进行了封装
}

func (op *OverseasProxy) distinguish(goods *Goods) bool {
	fmt.Println("对[", goods.Kind, "]进行了辨别真伪.")
	if goods.Fact == false {
		fmt.Println("发现假货", goods.Kind, ", 不应该购买。")
	}
	return goods.Fact
}

func (op *OverseasProxy) check(goods *Goods) {
	fmt.Println("对[", goods.Kind, "] 进行了海关检查， 成功的带回祖国")
}

func (op *OverseasProxy) Buy(goods *Goods) {
	if op.distinguish(goods) == true {
		//2. 进行购买
		op.shopping.Buy(goods) //调用原被代理的具体主题任务
		//3 海关安检
		op.check(goods)
	}
}

func NewProxy(shopping Shopping) Shopping {
	return &OverseasProxy{shopping}
}

// 代理模式更注重流程控制,在这个例子中就是在购买前后增加一些操作流程
// OverseasProxy和 KoreaShopping都实现了相同的接口：Shopping
func main() {
	g1 := Goods{
		Kind: "韩国面膜",
		Fact: true,
	}
	var shopping Shopping
	shopping = new(KoreaShopping) //具体的购买主题
	var overseasProxy Shopping
	overseasProxy = NewProxy(shopping)
	overseasProxy.Buy(&g1)
}
