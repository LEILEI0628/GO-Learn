package main

import "fmt"

func IfElse() {}
func IfNewVariable(start, end int) {
	if distance := end - start; distance > 100 { // distance作用域是整个IfElse代码段
		fmt.Println("远")
	} else if distance > 50 {
		fmt.Println("中")
	} else {
		fmt.Println("近")
	} // IfElse代码段结束后distance变量失效
}
