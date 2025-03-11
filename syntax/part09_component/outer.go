package main

import "fmt"

type Inner struct{}
type InnerV1 struct{}

func (i Inner) DoSomething() {
	fmt.Println("Inner DoSomething")
}

func (i InnerV1) DoSomething() {
	fmt.Println("InnerV1 DoSomething")
}

func (o Outer) DoSomething() {
	fmt.Println("Outer DoSomething")
}

type Outer struct {
	Inner // 可以组合多个
	InnerV1
}

type OuterPtr struct {
	*Inner // 不推荐使用此种方式（初始化OuterPtr后此处*Inner为nil）
}

func UseInner() {
	var out Outer // Outer会自动初始化Inner，OuterPtr不会
	// var out *Outer = &Outer{} // 指针也可以调用方法
	out.DoSomething() // 优先调用自己的
	// 当方法冲突时使用以下调用方式：
	out.Inner.DoSomething() // 没有(o Outer) DoSomething()时等价于out.DoSomething()
	var op OuterPtr         // 此时*Inner为nil，op.DoSomething()报空指针异常
	op = OuterPtr{
		Inner: &Inner{},
	}
	op.DoSomething()
}

func main() {
	UseInner()
}
