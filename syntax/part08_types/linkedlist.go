package main

type LinkedList struct {
	head *node
	tail *node
	Len  int // 包外可访问
}

func (l LinkedList) Add(index int, value any) {
	// (l LinkedList)是一个方法接收器，代表Add方法是定义在LinkedList上的

}

func (l *LinkedList) AddV1(index int, value any) {
	// (l *LinkedList)与(l LinkedList)不同

}

type node struct{}
