package LinkList

import "sort"

// 148. 排序链表 https://leetcode.cn/problems/sort-list/description/?envType=study-plan-v2&envId=top-100-liked
// 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

func sortList(head *ListNode) *ListNode {
	nodeMap := map[int][]*ListNode{} // 使用数组防止重复元素问题
	for head != nil {
		nodeMap[head.Val] = append(nodeMap[head.Val], head)
		head = head.Next
	}
	numList := make([]int, 0, len(nodeMap))
	for key, _ := range nodeMap {
		numList = append(numList, key)
	}
	sort.Ints(numList)
	result := &ListNode{}
	p := result
	for _, v := range numList {
		for _, node := range nodeMap[v] {
			p.Next = node
			p = p.Next
		}

	}
	p.Next = nil
	return result.Next
}
