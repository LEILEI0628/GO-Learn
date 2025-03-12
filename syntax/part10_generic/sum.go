package main

import "fmt"

func Sum[T Number](vals ...T) (sum T) {
	// 使用[T any]时无法计算加法，提示未定义“+”运算符
	// 引入泛型约束[T Number]
	for _, val := range vals {
		sum = sum + val
	}
	return
}

type Number interface {
	int | int64 | float64
}

func main() {
	fmt.Println(Sum(1, 2, 3, 4, 5))          // 默认解释为int类型
	fmt.Println(Sum[float64](1.1, 2.2, 3.3)) // 默认解释为float64类型，或指定自定义类型
}
