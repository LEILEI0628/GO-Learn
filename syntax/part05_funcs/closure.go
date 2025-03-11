package main

import "fmt"

func Closure(name string) func() string {
	return func() string { // 这个函数就是一个闭包
		// 返回的这个方法使用到了上下文“name”（参数），但是已经脱离了Closure方法
		return "Hello," + name
	}
}

func Closure1() func() int {
	var age int = 1
	fmt.Printf("OUT:%p\n", &age)
	return func() int { // 这个函数就是一个闭包
		age++
		fmt.Printf("IN:%p\n", &age)
		return age
	}
}

func main() {
	getAge1 := Closure1
	fmt.Println(getAge1()()) // 初始化了age变量
	fmt.Println(getAge1()())
	fmt.Println(getAge1()())

	getAge2 := Closure1()
	fmt.Println(getAge2()) // age变量仍在使用
	fmt.Println(getAge2())
	fmt.Println(getAge2())
}
