package main

import "fmt"

// Slice 切片
func Slice() {
	// 切片的底层是数组，切片相比于数组支持make创建、append元素、扩容等操作
	// 实际开发中尽量准确估计容量，防止扩容导致的资源浪费
	s1 := []int{1, 2, 3}
	fmt.Printf("%v, len=%d, cap=%d\n", s1, len(s1), cap(s1))
	s2 := make([]int, 2, 3)
	fmt.Printf("%v, len=%d, cap=%d\n", s2, len(s2), cap(s2))
	s3 := make([]int, 4) // 等价于make([]int, 4, 4)
	fmt.Printf("%v, len=%d, cap=%d\n", s3, len(s3), cap(s3))

	// 常规写法1：
	s4 := make([]int, 0, 4) // 后续使用append追加
	s4 = append(s4, 1)
	fmt.Printf("%v, len=%d, cap=%d\n", s4, len(s4), cap(s4))

	// 常规写法2（不常用）：
	s5 := make([]int, 4) // 类似数组的操作
	s5[0] = 1
	fmt.Printf("%v, len=%d, cap=%d\n", s5, len(s5), cap(s5))
}

func SubSlice() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:3] // 左闭右开，容量为[start:]（从start开始到父切片原本的底层数组的元素个数结束，即包括原数组后面几位的cap）
	fmt.Printf("%v, len=%d, cap=%d\n", s2, len(s2), cap(s2))
}

func main() {
	SubSlice()
}
