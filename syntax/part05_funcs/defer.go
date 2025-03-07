package main

import "fmt"

func Defer() {
	// 后进先出，类似于栈的结构，一般用于释放资源的场景
	defer func() {
		fmt.Println("First defer")
	}()

	defer func() {
		fmt.Println("Second defer")
	}()
}

func DeferClosure() int {
	i := 0
	defer func() {
		fmt.Println(i)
		i = 2
	}() // 在方法执行完成后才会执行（使用指针除外）
	i++
	return i
}

func DeferClosure1() (i int) {
	i = 0
	defer func(val int) {
		fmt.Println(val)
		i = 2 // 此处i类似于指针，更改成功
	}(i) // 值已经在此处传入
	i++
	return
}

func main() {
	fmt.Printf("return:%d\n", DeferClosure())
	fmt.Printf("return:%d\n", DeferClosure1())
}
