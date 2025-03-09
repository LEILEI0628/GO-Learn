package main

import "fmt"

// 衍生类型是一个全新的类型，类型定义与原类型相同但无法调用源类型的方法

type Integer int

func UseInt() {
	i1 := 10
	i2 := Integer(i1)
	var i3 Integer = 11
	fmt.Println(i2, i3)
}

type Fish struct {
	Name string
}

func (f Fish) FishName() {
	fmt.Println(f.Name)
}

func (f Fish) Swim() {
	fmt.Println("Fish Swim")
}

func (f FakeFish) Swim() {
	fmt.Println("FakeFish Swim")
}

type FakeFish Fish

// 类型别名：除了换名字其他没区别，常用于导出类型、兼容性修改等

type Yu = Fish

func UseFish() {
	f1 := Yu{Name: "Fish"} // f1就是Fish类型
	f1.FishName()
	f1.Swim()
	f2 := FakeFish(f1) // 因为结构定义相同可以转换，此处转换相当于复制，所以修改f2内容不会影响f1
	// f2.FishName() // f2没有FishName方法（衍生类型是一个全新的类型）
	f2.Swim()
	fmt.Println(f2.Name)
	f2.Name = "FakeFish"
	fmt.Println(f1.Name)
}

func main() {
	UseFish()
}
