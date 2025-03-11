package main

import "fmt"

// 在同一个包下不允许存在同名不同参数的方法
// 有返回值要保证一定返回（每一个分支）
// GO支持多个返回值
// 返回值可以命名（所有返回值都要命名）

// Func8 返回命名返回值（不使用名字的形式）
func Func8(age1 int) (name string, age int) {
	return "LEI LEI", age1
}

// Func9 返回命名返回值（使用名字的形式）
func Func9() (name string, age int) {
	func8 := Func8          // 方法可以作为变量，可以通过变量直接发起调用
	name1, age := func8(23) // 不需要的参数可以通过“_”忽略，使用:=时左侧必须要至少有一个新变量
	name = name1            // 返回值不能使用var变量声明和:=声明的形式，可以理解为返回值命名处已经声明过了
	//age = 23 // 未使用age时返回 "LEI LEI", 0，字符串返回""
	// 当返回值中有参数未被使用/赋值时，将返回默认值
	return // 等价于 return name, age
}

func Func10(age int, names ...string) {
	if len(names) > 0 {
		fmt.Printf("1 %v ", names)
		fmt.Println()
	}
	fmt.Printf("%d\n", age)
}

func Func11(age int, names ...any) {
	if len(names) > 0 {
		fmt.Printf("2 %v ", names)
	}
	fmt.Printf("%d\n", age)
}

func main() {
	name := []string{"A", "B", "C"}
	Func10(18, name...) // name...表示解切片，传入的其实是三个string
	Func11(18, name)    // 传入的其实是一个[]string
}
