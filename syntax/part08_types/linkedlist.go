package main

import "time"

type LinkedList struct {
	head *node
	tail *node
	Len  int // 包外可访问

	CreateTime time.Time
}

func (l LinkedList) Add(index int, value any) {
	// (l LinkedList)是一个方法接收器，代表Add方法是定义在LinkedList上的

}

func (l *LinkedList) AddV1(index int, value any) {
	// (l *LinkedList)与(l LinkedList)不同

}

type node struct {
	// 自引用不用指针会编译错误（无法计算大小）
	prev *node
	next *node
}
