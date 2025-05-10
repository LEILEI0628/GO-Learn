package main

// 21. 合并两个有序链表 https://leetcode.cn/problems/merge-two-sorted-lists/description/?envType=study-plan-v2&envId=top-100-liked
// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

// 方法一：迭代
// 思路
// 我们可以用迭代的方法来实现上述算法。当 l1 和 l2 都不是空链表时，判断 l1 和 l2 哪一个链表的头节点的值更小，将较小值的节点添加到结果里，当一个节点被添加到结果里之后，将对应链表中的节点向后移一位。
// 算法
// 首先，我们设定一个哨兵节点 prehead ，这可以在最后让我们比较容易地返回合并后的链表。我们维护一个 prev 指针，我们需要做的是调整它的 next 指针。然后，我们重复以下过程，直到 l1 或者 l2 指向了 null ：如果 l1 当前节点的值小于等于 l2 ，我们就把 l1 当前的节点接在 prev 节点的后面同时将 l1 指针往后移一位。否则，我们对 l2 做同样的操作。不管我们将哪一个元素接在了后面，我们都需要把 prev 向后移一位。
// 在循环终止的时候， l1 和 l2 至多有一个是非空的。由于输入的两个链表都是有序的，所以不管哪个链表是非空的，它包含的所有元素都比前面已经合并链表中的所有元素都要大。这意味着我们只需要简单地将非空链表接在合并链表的后面，并返回合并链表即可。

func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	p := &ListNode{}
	list := p
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p = p2
			p2 = p2.Next
		}
	}
	if p1 == nil {
		p.Next = p2
	} else {
		p.Next = p1
	}
	return list.Next
}

// 思路
// 直接用 mergeTwoLists 当作递归函数：
// 递归边界：如果其中一个链表为空，直接返回另一个链表作为合并后的结果。
// 如果两个链表都不为空，则比较两个链表当前节点的值，并选择较小的节点作为新链表的当前节点。例如 list1的节点值更小，那么递归调用 mergeTwoLists(list1.next, list2)，将递归返回的链表接在 list1的末尾。

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2 // 注：如果都为空则返回空
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}
