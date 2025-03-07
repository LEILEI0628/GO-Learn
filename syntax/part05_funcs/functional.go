package main

func Functional1() {
	println("Functional1")
}

func UseFunctional1() {
	myFunc := Functional1 // Functional1无需加“()”
	myFunc()
}
