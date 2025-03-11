package main // main方法必须要在main包运行

import (
	"fmt"
	"unicode/utf8"
)

func main() { // 无参数，无返回值
	fmt.Println("Hello GO!")
	fmt.Println(`HELLO
GO!`) // ``适用大段文字，可以换行
	fmt.Println("\"Hello GO!\"")              // 转义
	fmt.Println(len("你好"))                    // 计算字节长度
	fmt.Println(utf8.RuneCountInString("你好")) // 计算字符个数，不推荐使用len()除字符编码字节长度
	a := make([]int, 3, 4)
	a = append(a, 1)
	fmt.Printf("%v %d\n", a, a[3])
}
