package LinkList

import "testing"

func TestMyLinkedList(t *testing.T) {
	// Your MyLinkedList object will be instantiated and called as such:
	obj := Constructor()
	obj.AddAtHead(2)
	obj.AddAtHead(1)
	obj.AddAtTail(4)
	obj.AddAtIndex(3, 3)
	obj.DeleteAtIndex(3)
	t.Log(obj.Get(3))
	t.Log(printList(obj))

}

func printList(head MyLinkedList) (result []int) {
	cur := head.Next
	for cur != nil {
		result = append(result, cur.Val)
		cur = cur.Next
	}
	return
}
