package main

// 206. 反转链表 https://leetcode.cn/problems/reverse-linked-list/description/?envType=study-plan-v2&envId=top-100-liked
// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p, q := head, head.Next
	head.Next = nil
	for q != nil {
		temp := q.Next
		q.Next = p
		p = q
		q = temp
	}
	return p
}

func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
