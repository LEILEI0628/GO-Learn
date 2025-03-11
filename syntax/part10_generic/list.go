package main

type List[T any] interface { // T：类型参数，约束是any
	Add(index int, value T)
	Append(value T)
}

func UseList() {
	//var l List[int]
	//l.Append(10)
}

func main() {
	UseList()
}
