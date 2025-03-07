package main

import "fmt"

func ForLoop() {
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	for i := 0; i < 3; { // 第三个条件可以没有
		if i == 1 {
			continue // GO语言中也有continue和break关键字
		}
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()

	// 类似while的写法
	i := 0
	for i < 3 {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()

	// 死循环（慎用！）
	for { // 或for true
		fmt.Println("死循环")
		break
	}
}

func ForRange() {
	// 数组/切片
	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		fmt.Printf("%d:%d ", i, v)
	}
	fmt.Println()

	for i := range arr { // 省略index写法：_, v := range arr
		fmt.Printf("%d:%d ", i, arr[i])
	}

	// Map
	m := map[int]string{1: "a", 2: "b", 3: "c"} // Map的顺序是随机的
	for k, v := range m {                       // 省略写法同上（Map可以通过kay索引）
		fmt.Printf("%d:%s ", k, v)
	}
	fmt.Println()

	// 不要对迭代参数取地址！
}
