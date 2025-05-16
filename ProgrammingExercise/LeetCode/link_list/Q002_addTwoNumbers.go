package LinkList

// 2. 两数相加 https://leetcode.cn/problems/add-two-numbers/description/?envType=study-plan-v2&envId=top-100-liked
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	promotion := 0 // 进位
	head := &ListNode{}
	m := head
	for promotion != 0 || l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			num := l1.Val + l2.Val + promotion
			promotion = num / 10
			m.Next = &ListNode{Val: num % 10}
			m = m.Next
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			num := l1.Val + promotion
			promotion = num / 10
			m.Next = &ListNode{Val: num % 10}
			m = m.Next
			l1 = l1.Next
		} else if l2 != nil {
			num := l2.Val + promotion
			promotion = num / 10
			m.Next = &ListNode{Val: num % 10}
			m = m.Next
			l2 = l2.Next
		} else {
			m.Next = &ListNode{Val: promotion}
			break
		}
	}
	return head.Next
}
