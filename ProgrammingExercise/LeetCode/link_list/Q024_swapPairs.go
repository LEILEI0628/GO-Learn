package main

// 24. 两两交换链表中的节点 https://leetcode.cn/problems/swap-nodes-in-pairs/description/?envType=study-plan-v2&envId=top-100-liked
// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

func swapPairs(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	result := node.Next
	p := swapPairs(node.Next.Next)
	node.Next.Next = node
	node.Next = p

	return result
}

// 数组
func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nodeList := make([]*ListNode, 0)
	p := head
	for p != nil {
		nodeList = append(nodeList, p)
		p = p.Next
	}
	last := len(nodeList)
	last -= last % 2
	for i := last - 1; i >= 0; i -= 2 {
		var temp *ListNode          // 初始值为nil，若为最后两个节点则temp为nil
		if i+1 == len(nodeList)-1 { // 节点数为奇数的情况
			temp = nodeList[i+1]
		} else if i+2 < len(nodeList) { // 非最后两个节点的情况
			temp = nodeList[i+2]
		}
		nodeList[i].Next = nodeList[i-1]
		nodeList[i-1].Next = temp
	}
	return nodeList[1]
}

// 数组2
func swapPairs3(head *ListNode) *ListNode {
	nodeList := make([]*ListNode, 0)
	p := head
	for p != nil {
		nodeList = append(nodeList, p)
		p = p.Next
	}
	result := &ListNode{}
	q := result
	for i := 0; i < len(nodeList); i += 2 {
		if i+1 == len(nodeList) { // 最后仅剩一个元素的情况
			q.Next = nodeList[i]
			q = q.Next
			break
		}
		q.Next = nodeList[i+1]
		nodeList[i+1].Next = nodeList[i]
		q = nodeList[i]
	}
	q.Next = nil
	return result.Next
}

// 数组3
func swapPairs4(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nodeList := make([]*ListNode, 0)
	p := head
	for p != nil {
		nodeList = append(nodeList, p)
		p = p.Next
	}
	nodeList[1].Next = nodeList[0]
	p = nodeList[0]
	for i := 2; i < len(nodeList); i += 2 {
		if i+1 == len(nodeList) { // 最后仅剩一个元素的情况
			p.Next = nodeList[i]
			p = p.Next
			break
		}
		p.Next = nodeList[i+1]
		nodeList[i+1].Next = nodeList[i]
		p = nodeList[i]
	}
	p.Next = nil
	return nodeList[1]
}

func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	result := head.Next
	p := head
	for p != nil {
		if p.Next.Next != nil && p.Next.Next.Next == nil { // 奇数最后三个节点的情况
			temp := p.Next.Next
			p.Next.Next = p
			p.Next = temp
			break
		} else if p.Next.Next == nil { // 偶数最后两个节点的情况
			p.Next.Next = p
			p.Next = nil
			break
		}
		next := p.Next.Next      // 下一个p节点的位置
		temp := p.Next.Next.Next // 当前节点指向的更改后的下一个节点的位置
		p.Next.Next = p
		p.Next = temp
		p = next
	}
	return result
}
