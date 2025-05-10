package main

// 234. 回文链表 https://leetcode.cn/problems/palindrome-linked-list/description/?envType=study-plan-v2&envId=top-100-liked

// 方法一：将值复制到数组中后用双指针法
func isPalindrome(head *ListNode) bool {
	list := make([]int, 0)
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	for i := 0; i < len(list)/2; i++ {
		if list[i] != list[len(list)-i-1] {
			return false
		}
	}
	return true
}

// 递归
func isPalindrome1(head *ListNode) bool {
	frontPointer := head
	var recursivelyCheck func(*ListNode) bool
	recursivelyCheck = func(curNode *ListNode) bool {
		if curNode != nil {
			if !recursivelyCheck(curNode.Next) {
				return false
			}
			if curNode.Val != frontPointer.Val {
				return false
			}
			frontPointer = frontPointer.Next
		}
		return true
	}
	return recursivelyCheck(head)
}

// 递归
var pre *ListNode

func isPalindrome2(head *ListNode) bool {
	pre = head
	return f1(head)
}

func f1(node *ListNode) bool {
	if node != nil {
		if !f1(node.Next) {
			return false
		}
		if node.Val != pre.Val {
			pre = pre.Next
			return false
		}
		pre = pre.Next
	}
	return true
}
