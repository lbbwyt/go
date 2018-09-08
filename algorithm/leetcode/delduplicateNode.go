package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}
	current := head
	for current != nil && current.next != nil {
		if current.v == current.next.v {
			current.next = current.next.next
		} else {
			current = current.next
		}
	}
	return head
}

func main() {

}

//Linkedlist链表
type ListNode struct {
	v    int
	next *ListNode
}
