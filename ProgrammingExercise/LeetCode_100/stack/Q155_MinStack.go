package main

// 155. 最小栈 https://leetcode.cn/problems/min-stack/description/?envType=study-plan-v2&envId=top-100-liked
// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
// 实现 MinStack 类:
// MinStack() 初始化堆栈对象。
// void push(int val) 将元素val推入堆栈。
// void pop() 删除堆栈顶部的元素。
// int top() 获取堆栈顶部的元素。
// int getMin() 获取堆栈中的最小元素。

type MinStack struct {
}

func Constructor() MinStack {
	return MinStack{}
}

func (stack *MinStack) Push(val int) {
}

func (stack *MinStack) Pop() {

}

func (stack *MinStack) Top() int {
	return 0
}

func (stack *MinStack) GetMin() int {
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
