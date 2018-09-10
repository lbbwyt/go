package main

//Linkedlist链表
type Node struct {
	v    int
	next *Node
}

func reverse(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	var p, q, pr *Node

	//首先让头节点与第一个元素节点断开，但是要注意在断开之前需要用p指针指向第一个元素节点来保存第一个元素节点的位置，然后再断开。
	//在这里有一个指针q指向一个指针域为空的节点，这个节点用来做为链表反转后的最后一个节点。
	p = head.next
	head.next = nil
	//q始终指向反转后链表中最靠近头结点的一个节点
	q = nil
	for p != nil {
		//保存p指针
		pr = p.next
		//头插法将p连接到q的后面，并使q指向最新节点
		p.next = q //使p的next指向q
		q = p
		//恢复p指针，并指向下一个节点
		p = pr
	}

	head.next = q
	return head
}

func main() {

}
