package main

func Functional1() {
	fmt.Println("Functional1")
}

func UseFunctional1() {
	myFunc := Functional1 // Functional1无需加“()”
	myFunc()
}
