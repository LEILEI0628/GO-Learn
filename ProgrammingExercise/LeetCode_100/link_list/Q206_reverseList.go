package main

// 206. 反转链表 https://leetcode.cn/problems/reverse-linked-list/description/?envType=study-plan-v2&envId=top-100-liked
// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

func reverseList(head *ListNode) *ListNode {
	if head == nil { // 传进来的是个空指针
		return head
	}
	if head.Next == nil { // 到达最后一个节点
		return head
	}
	newHead := reverseList(head.Next) // 1.递归修改剩下节点
	head.Next.Next = head             // 2.将当前节点的下一个节点的Next指针指向当前节点
	head.Next = nil                   // 3.将当前节点的Next指针置空
	return newHead
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var p *ListNode
	q := head
	for q != nil {
		temp := q.Next
		q.Next = p
		p = q
		q = temp
	}
	return p
}
