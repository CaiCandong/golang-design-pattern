package main

import "fmt"

// ======= 抽象层 =========
// 接口
type AbstractGPU interface {
	Display()
}
type AbstractCPU interface {
	Calculate()
}
type AbstractMemory interface {
	Storage()
}

//抽象工厂
type AbstractComputerFactory interface {
	CreateGPU() AbstractGPU
	CreateCPU() AbstractCPU
	CreateMemory() AbstractMemory
}

// ======= 基础类模块 =========
type NvidiaCPU struct{}

func (cpu *NvidiaCPU) Calculate() {
	fmt.Println("calculate by  NvidiaCPU ")
}

type NvidiaGPU struct{}

func (gpu *NvidiaGPU) Display() {
	fmt.Println("display by  NvidiaGPU ")
}

type NvidiaMemory struct{}

func (m *NvidiaMemory) Storage() {
	fmt.Println("storage by NvidiaMemory ")
}

type IntelCPU struct{}

func (cpu *IntelCPU) Calculate() {
	fmt.Println("calculate by  IntelCPU")
}

type IntelGPU struct{}

func (gpu *IntelGPU) Display() {
	fmt.Println("display by  IntelGPU")
}

type IntelMemory struct{}

func (m *IntelMemory) Storage() {
	fmt.Println("storage by  IntelMemory ")
}

type KingstonCPU struct{}

func (cpu *KingstonCPU) Calculate() {
	fmt.Println("calculate by KingstonCPU ")
}

type KingstonGPU struct{}

func (gpu *KingstonGPU) Display() {
	fmt.Println("display by  KingstonGPU")
}

type KingstonMemory struct{}

func (m *KingstonMemory) Storage() {
	fmt.Println("storage by KingstonMemory ")
}

// Intel工厂
type IntelFactory struct{}

func (factory *IntelFactory) CreateGPU() AbstractGPU {
	var product AbstractGPU
	product = new(IntelGPU)
	return product
}
func (factory *IntelFactory) CreateCPU() AbstractCPU {
	var product AbstractCPU
	product = new(IntelCPU)
	return product
}
func (factory *IntelFactory) CreateMemory() AbstractMemory {
	var product AbstractMemory
	product = new(IntelMemory)
	return product
}

// Nvidia工厂
type NvidiaFactory struct{}

func (factory *NvidiaFactory) CreateGPU() AbstractGPU {
	var product AbstractGPU
	product = new(NvidiaGPU)
	return product
}
func (factory *NvidiaFactory) CreateCPU() AbstractCPU {
	var product AbstractCPU
	product = new(NvidiaCPU)
	return product
}
func (factory *NvidiaFactory) CreateMemory() AbstractMemory {
	var product AbstractMemory
	product = new(NvidiaMemory)
	return product
}

// Kingston工厂
type KingstonFactory struct{}

func (factory *KingstonFactory) CreateGPU() AbstractGPU {
	var product AbstractGPU
	product = new(KingstonGPU)
	return product
}
func (factory *KingstonFactory) CreateCPU() AbstractCPU {
	var product AbstractCPU
	product = new(KingstonCPU)
	return product
}
func (factory *KingstonFactory) CreateMemory() AbstractMemory {
	var product AbstractMemory
	product = new(KingstonMemory)
	return product
}

// ======== 业务逻辑层 =======
type Computer struct {
	CPU    AbstractCPU
	GPU    AbstractGPU
	Memory AbstractMemory
}

func (c *Computer) Stats() {
	fmt.Println("-------------computer stats---------------")
	c.CPU.Calculate()
	c.GPU.Display()
	c.Memory.Storage()
}

func main() {

	intel_factory := new(IntelFactory)
	nvidia_factor := new(NvidiaFactory)
	kingston_factory := new(KingstonFactory)

	computer1 := Computer{
		GPU:    intel_factory.CreateGPU(),
		CPU:    intel_factory.CreateCPU(),
		Memory: intel_factory.CreateMemory(),
	}
	computer2 := Computer{
		GPU:    nvidia_factor.CreateGPU(),
		CPU:    intel_factory.CreateCPU(),
		Memory: kingston_factory.CreateMemory(),
	}
	computer1.Stats()
	computer2.Stats()

}
