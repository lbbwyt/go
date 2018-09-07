package main

import "fmt"

func main() {
	head := &node{
		v:    0,
		next: nil,
	}
	arr := []int{7, 5, 3, 4}
	newhead := CreateLink(head, arr)
	printLink(newhead)

	reversenode := reverse(newhead)
	printLink(reversenode)

}

//Linkedlist链表
type node struct {
	v    int
	next *node
}

//初始化一个长度为length(不包括头节点)的链表，节点的值默认为0,
//尾插法
func CreateLink(head *node, arr []int) *node {
	length := len(arr)
	if length <= 0 {
		return head
	}
	for i := length - 1; i >= 0; i-- {
		p := &node{arr[i], nil}
		p.next = head.next
		head.next = p
	}
	return head
}

func reverse(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}
	var reversedHead *node
	var p *node

	reversedHead = head
	head = head.next
	reversedHead.next = nil
	p = head.next
	for head != nil {
		head.next = reversedHead
		reversedHead = head
		head = p
		if p != nil {
			p = p.next
		}
	}
	return reversedHead
}

func reverseLink(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}
	var reverseNode, p *node
	reverseNode = head
	//如果没有head = head.next,那么执行reverseNode.next = nil之后，head.next也为nil
	head = head.next
	reverseNode.next = nil
	p = head
	//头插法实现链表反转
	for p != nil {
		temp := &node{p.v, nil}
		temp.next = reverseNode.next
		reverseNode.next = temp

		p = p.next

	}
	return reverseNode
}

func printLink(head *node) {
	for p := head; p != nil; p = p.next {
		fmt.Print(p.v)
		fmt.Print(" ")
	}
	fmt.Println()
}
