package LinkList

// 203. 移除链表元素 https://leetcode.cn/problems/remove-linked-list-elements/description/
// 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
func removeElements(head *ListNode, val int) *ListNode { // 使用额外的头结点
	result := &ListNode{Next: head}     // 额外的头节点
	cur := result                       // p指向需要删除的前一个元素
	for cur != nil && cur.Next != nil { // cur指向需要删除的前一个元素（删除的是cur的下一个元素）
		if cur.Next.Val == val { // 删除元素
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return result.Next
}

func removeElementsV1(head *ListNode, val int) *ListNode { // 不使用额外的头结点（需要单独处理第一个节点）
	for head != nil && head.Val == val {
		head = head.Next
	}
	cur := head
	for cur != nil && cur.Next != nil { // cur指向需要删除的前一个元素（删除的是cur的下一个元素）
		if cur.Next.Val == val { // 删除元素
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
