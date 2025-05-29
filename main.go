package main // main方法必须要在main包运行

import (
	"fmt"
)

func main() { // 无参数，无返回值
	arr := []int{10, 7, 8, 9, 1, 5}
	quick(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quick(arr []int, low int, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quick(arr, low, pi-1)
		quick(arr, pi+1, high)
	}
}

func partition(arr []int, low int, high int) int {
	pi := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pi {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
