package main

import "fmt"

func NewUser() {
	// 初始化结构体
	u := User{} // u是结构体
	fmt.Printf("%#v\n", u)

	up := &User{} // up是指针
	//up := new(User) // 不推荐使用此种方式
	// 指针也可以调用接收器方法，GO会自动解引用
	up.changeName("BBB")
	up.changeAge(18)
	fmt.Printf("%#v\n", up)
	fmt.Printf("%#v\n", *up) // 解引用

	// 初始化并赋值
	u1 := User{Name: "AAA", Age: 18}
	// u1 := User{"AAA", 18} // 不推荐使用这种方式
	u1.Age = 20
	fmt.Printf("%#v\n", u1)

	var up1 *User
	fmt.Printf("%#v\n", up1)
	// fmt.Println(up1.Name) // 空指针异常
}

func ChangeUser() {
	user := User{Name: "AAA", Age: 18}
	fmt.Printf("user:%p\n", &user)
	// 需要初始化user后才能调用changeName等方法
	user.changeName("BBB")
	user.changeAge(30)

}

func (user User) changeName(name string) {
	// (user User)是一个方法接收器，代表ChangeName方法是定义在User上的
	user.Name = name
	fmt.Printf("userChangeName:%p\n", &user)
}

func (user *User) changeAge(age int) {
	//值传递问题：(user *User)与(user User)不同，
	//(user *User)传入的是user的指针，所以可以对user进行更改。
	// 而(user User)传入的是结构体（复制了一个user），所以无法对原user进行修改。
	user.Age = age
	fmt.Printf("userChangeAge:%p\n", user)
}

type User struct {
	Name string
	Age  int
}
