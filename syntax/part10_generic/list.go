package main

import "fmt"

type List[T any] interface { // T：类型参数，约束是any
	Add(index int, value T)
	Append(value T)
}

func UseList() {
	//var l List[int] // 空指针，没有初始化，只是一个接口
	//l.Append(10)

	lk := LinkedList[int]{}
	intVal := lk.head.val
	fmt.Println(intVal)

}

type LinkedList[T any] struct {
	head *node[T]
	t    T
}

type node[T any] struct {
	val T
}

func main() {
	UseList()
}
