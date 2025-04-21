package main

import "math"

// 2. 两数相加 https://leetcode.cn/problems/add-two-numbers/description/?envType=study-plan-v2&envId=top-100-liked
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var num1 int64
	var num2 int64
	for i := 0; l1 != nil; i++ {
		num1 += int64(l1.Val) * int64(math.Pow(10, float64(i)))
		l1 = l1.Next
	}
	for i := 0; l2 != nil; i++ {
		num2 += int64(l2.Val) * int64(math.Pow(10, float64(i)))
		l2 = l2.Next
	}
	num1 += num2
	nums := int64(num1)
	if nums == 0 {
		return &ListNode{Val: 0}
	}
	p := &ListNode{}
	result := p
	for nums != 0 {
		p.Next = &ListNode{Val: int(nums % 10)}
		p = p.Next
		nums = nums / 10
	}
	return result.Next
}
