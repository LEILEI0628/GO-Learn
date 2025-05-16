package LinkList

// 707. 设计链表 https://leetcode.cn/problems/design-linked-list/description/
// 你可以选择使用单链表或者双链表，设计并实现自己的链表。
// 单链表中的节点应该具备两个属性：val 和 next 。val 是当前节点的值，next 是指向下一个节点的指针/引用。
// 如果是双向链表，则还需要属性 prev 以指示链表中的上一个节点。假设链表中的所有节点下标从 0 开始。
// 实现 MyLinkedList 类：
// MyLinkedList() 初始化 MyLinkedList 对象。
// int get(int index) 获取链表中下标为 index 的节点的值。如果下标无效，则返回 -1 。
// void addAtHead(int val) 将一个值为 val 的节点插入到链表中第一个元素之前。在插入完成后，新节点会成为链表的第一个节点。
// void addAtTail(int val) 将一个值为 val 的节点追加到链表中作为链表的最后一个元素。
// void addAtIndex(int index, int val) 将一个值为 val 的节点插入到链表中下标为 index 的节点之前。如果 index 等于链表的长度，那么该节点会被追加到链表的末尾。如果 index 比长度更大，该节点将 不会插入 到链表中。
// void deleteAtIndex(int index) 如果下标有效，则删除链表中下标为 index 的节点。

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}
	cur := l
	for i := 0; i <= index; i++ {
		cur = cur.Next
		if cur == nil {
			return -1
		}
	}
	return cur.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.Next = &MyLinkedList{Val: val, Next: l.Next}
}

func (l *MyLinkedList) AddAtTail(val int) {
	cur := l
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &MyLinkedList{Val: val}
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	cur := l
	for i := 0; i < index-1; i++ { // 指向插入位置的前一个节点
		cur = cur.Next
		if cur == nil {
			return
		}

	}
	if cur.Next != nil {
		node := &MyLinkedList{Val: val, Next: cur.Next}
		cur.Next = node
	} else { // 插入位置为最后
		cur.Next = &MyLinkedList{Val: val}
	}

}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}
	cur := l
	for i := 0; i < index; i++ {
		cur = cur.Next
		if cur.Next == nil {
			return
		}
	}
	cur.Next = cur.Next.Next
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
