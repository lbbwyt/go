// linklist
package algorithm

import (
	"fmt"
)

//Linkedlist链表

type Node struct {
	v    int
	next *Node
}

var head *Node = nil

//定义Node上的方法
func (l *Node) PushBack(val int) *Node {
	/* 没有节点时头指针指向l */
	if head == nil {
		l.v = val
		l.next = nil
		head = l
		return l
	} else {
		//		遍历到插入节点位置的前一个节点
		for l.next != nil {
			l = l.next
		}
		l.next = new(Node)
		l.next.v = val
		l.next.next = nil
		return l
	}
}

func (l *Node) PopBack() *Node {
	//为空的处理
	if head == nil {
		return head
	}
	//只有一个元素的处理，head是指向第一个元素的
	if head.next == nil {
		head = nil
		return head
	}
	cpnode := new(Node)
	cpnode = head
	for cpnode.next.next != nil {
		cpnode = cpnode.next
	}
	cpnode.next = nil
	return head
}

func (l *Node) Print() {
	if head == nil {
		fmt.Println("empty")
	} else {
		cpnode := head
		for {
			fmt.Println(cpnode.v)
			if cpnode.next == nil {
				break
			}
			cpnode = cpnode.next

		}
	}

}
