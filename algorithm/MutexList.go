package main

import (
	"fmt"
	"sync"
)

//实现一个并发安全的链表

type Node struct {
	Value interface{}
	next *Node
}

func NewNode(v interface{}) *Node {
	n:= &Node{Value:v}
	n.next = nil
	return  n
}

type MutexList struct {
	locker *sync.Mutex
	head, tail *Node //head是一个哨兵节点，不存储实际的数值。
	size int64
}
//尾部插入
func (l *MutexList) PushBack(v interface{}) bool {
	if v == nil {
		return  false
	}
	l.locker.Lock()
	defer l.locker.Unlock()
	node := NewNode(v)
	l.tail.next = node
	l.tail = node
	l.size ++
	return  true
}


func (l *MutexList) PushFront(v interface{}) bool {
	if v == nil {
		return  false
	}
	l.locker.Lock()
	defer l.locker.Unlock()
	node := NewNode(v)
    p := l.head.next
    l.head.next = node
    node.next = p
    l.size ++
    if l.size ==1 {
    	l.tail = node
	}
    return true
}

func (l *MutexList) print() {
	curNode := l.head
	p := curNode.next

	for p!=nil {
		fmt.Println(p.Value)
		p=p.next
	}
}



//交替打印数字和字母
func main() {
head := NewNode("")
muList := &MutexList{
	locker:&sync.Mutex{},
	head:head,
	tail:head,
	size:0,
}
	muList.PushBack(3)
	muList.PushFront(2)
	muList.PushBack(4)
    muList.print()

}

