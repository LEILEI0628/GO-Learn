package main

import "fmt"

func Switch(level int) {
	switch level {
	case 1:
		fmt.Println("1级")
		break // go语言中可以不加break
	case 2:
		fmt.Println("2级")
	case 3:
		fmt.Println("3级")
	default:
		fmt.Println("？级")
	}

	switch {
	case level == 1: // 使用这种写法时尽量保证每个判定条件都是互斥的，代码可读性更高
		fmt.Println("1级")
	case level == 2:
		fmt.Println("2级")
	case level == 3:
		fmt.Println("3级")
	default:
		fmt.Println("？级")
	}
}

func main() {
	Switch(3)
}
