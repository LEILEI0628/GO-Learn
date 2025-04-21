package main

// 160. 相交链表 https://leetcode.cn/problems/intersection-of-two-linked-lists/description/?envType=study-plan-v2&envId=top-100-liked
// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。

// 方法一：哈希集合（写法2）
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	linkAMap := map[*ListNode]bool{}
	linkA := headA
	for linkA != nil {
		linkAMap[linkA] = true
		linkA = linkA.Next
	}

	linkB := headB
	for linkB != nil {
		if linkAMap[linkB] {
			return linkB
		}
		linkB = linkB.Next
	}
	return nil
}

// 方法一：哈希集合
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}

// 方法二：双指针
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb { // pa pb最多移动两个链表长度之和，此时均为空；或当pa=pb时链表相交
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
