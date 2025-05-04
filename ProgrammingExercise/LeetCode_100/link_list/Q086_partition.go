package main

// 86. 分隔链表 https://leetcode.cn/problems/partition-list/
// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
// 你应当 保留 两个分区中每个节点的初始相对位置。
func partition(head *ListNode, x int) *ListNode {
	left := ListNode{}
	p := &left
	right := ListNode{}
	q := &right
	for head != nil { // 对链表遍历
		if head.Val < x { // 小于x的挂到左边
			p.Next = head
			p = head
		} else { // 大于x的挂到右边
			q.Next = head
			q = head
		}
		head = head.Next
	}
	p.Next = right.Next // 拼接链表
	q.Next = nil
	return left.Next
}
