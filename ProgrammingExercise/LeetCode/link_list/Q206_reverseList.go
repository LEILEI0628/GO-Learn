package LinkList

// 206. 反转链表 https://leetcode.cn/problems/reverse-linked-list/description/?envType=study-plan-v2&envId=top-100-liked
// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

func reverseList(head *ListNode) *ListNode {
	var reverse func(pre *ListNode, cur *ListNode) *ListNode
	reverse = func(pre *ListNode, cur *ListNode) *ListNode {
		if cur == nil {
			return pre // 返回翻转后的头结点
		}
		head := reverse(cur, cur.Next) // 修改后续的节点
		cur.Next = pre                 // 修改当前节点
		return head                    // 逐级返回翻转后的头结点
	}
	return reverse(nil, head)
}

func reverseListV1(head *ListNode) *ListNode { // 递归写法
	var result *ListNode
	var reverse func(pre *ListNode, cur *ListNode)
	reverse = func(pre *ListNode, cur *ListNode) {
		if cur == nil {
			result = pre
			return
		}
		reverse(cur, cur.Next) // 修改后续的节点
		cur.Next = pre         // 修改当前节点
	}
	reverse(nil, head)
	return result
}

func reverseListV2(head *ListNode) *ListNode { // 双指针写法
	var preNode *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = preNode
		preNode = cur
		cur = tmp
	}
	return preNode
}

func reverseListV3(cur *ListNode) *ListNode { // 双指针简化写法
	var pre *ListNode
	for cur != nil {
		cur, cur.Next, pre = cur.Next, pre, cur
	}
	return pre
}
