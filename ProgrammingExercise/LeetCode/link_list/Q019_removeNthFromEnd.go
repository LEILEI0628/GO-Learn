package LinkList

// 19. 删除链表的倒数第 N 个结点 https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/?envType=study-plan-v2&envId=top-100-liked
// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

// 方法一：计算链表长度
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	index := 0
	p := head
	for p != nil {
		index++
		p = p.Next
	}
	index -= n
	temp := &ListNode{Next: head}
	q := temp
	for i := 0; i < index+1; i++ {
		if i == index {
			q.Next = q.Next.Next
		}
		q = q.Next
	}
	return temp.Next
}

// 方法三：双指针
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	result := &ListNode{Next: head}
	slow := result
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return result.Next
}

// 数组
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	list := make([]*ListNode, 0)
	temp := &ListNode{Next: head}
	p := temp
	for p != nil {
		list = append(list, p)
		p = p.Next
	}
	list[len(list)-n-1].Next = list[len(list)-n].Next
	return temp.Next
}

// 方法二：栈
// 思路与算法
// 我们也可以在遍历链表的同时将所有节点依次入栈。根据栈「先进后出」的原则，我们弹出栈的第 n 个节点就是需要删除的节点，并且目前栈顶的节点就是待删除节点的前驱节点。这样一来，删除操作就变得十分方便了。

func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}
