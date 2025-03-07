package main

import "fmt"

// Recursive 打印从1到i
func Recursive(i int) {
	// 递归使用不当会出现stack overflow
	if i == 0 {
		return
	}
	Recursive(i - 1)
	fmt.Printf("%d ", i)
}
