package LinkList

// 25. K 个一组翻转链表 https://leetcode.cn/problems/reverse-nodes-in-k-group/description/?envType=study-plan-v2&envId=top-100-liked

func reverseKGroup(head *ListNode, k int) *ListNode {
	nodeList := make([]*ListNode, 0)
	p := head
	for p != nil {
		nodeList = append(nodeList, p)
		p = p.Next
	}
	head = &ListNode{}
	q := head
	for i := 0; i < len(nodeList); i += k {
		if i+k-1 >= len(nodeList) { // 最后的的节点不够交换
			q.Next = nodeList[i]
			break
		}
		for j := k - 1; j >= 0; j-- {
			q.Next = nodeList[i+j]
			nodeList[i+j].Next = nil // .Next不置为nil可能会导致环
			q = nodeList[i+j]
		}
	}
	return head.Next
}

// 方法一：模拟
// 思路与算法
// 本题的目标非常清晰易懂，不涉及复杂的算法，但是实现过程中需要考虑的细节比较多，容易写出冗长的代码。主要考查面试者设计的能力。
func reverseKGroup1(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = myReverse(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}
