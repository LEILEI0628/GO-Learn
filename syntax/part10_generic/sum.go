package main

import (
	"encoding/json"
	"fmt"
)

func Sum[T Number](vals ...T) (sum T) {
	// 泛型约束不能用在参数和返回值中，即Sum[T Number](vals ...Number) (sum Number)是错误写法
	// 使用[T any]时无法计算加法，提示未定义“+”运算符
	// 引入泛型约束[T Number]
	for _, val := range vals {
		sum = sum + val
	}
	return
}

type Number interface {
	~int | int64 | float64 // ~int：int及其衍生类型
}

type Integer int

func main() {
	fmt.Println(Sum(1, 2, 3, 4, 5)) // 默认解释为int类型
	// 指定的自定义类型需要是泛型约束中包含的
	fmt.Println(Sum[Integer](1, 2, 3, 4, 5)) // 默认解释为int类型
	fmt.Println(Sum[float64](1.1, 2.2, 3.3)) // 默认解释为float64类型，或指定自定义类型
}

func ReleaseResource[R json.Marshaler](r R) { // [R json.Marshaler]：可以传入接口或any（其他类型不支持）
	r.MarshalJSON()
}
