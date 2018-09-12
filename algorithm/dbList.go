package main

import "fmt"

//节点定义， 和list的element一致
type Node struct {
	 next, prev *Node
	 list *dbList
	 Value interface{}
}

//双向链表
type dbList struct {
	 head Node //头结点
	 len int //链表长度
}

//初始化一个空的双向链表，
func (l *dbList) Init() *dbList {
	l.head.next = &l.head
	l.head.prev = &l.head
	l.len = 0
	return  l
}

// .
func (l *dbList) lazyInit() {
	if l.head.next == nil {
		l.Init()
	}
}


//构造函数
func New() *dbList{
	return  new(dbList).Init()
}
//将e插入到at后，返回e
func(l *dbList) insert(e,  at *Node) *Node {
 n := at.next
 at.next = e
 e.prev = at
 e.next = n
 n.prev = e
 e.list = l
 l.len ++
 return  e
}


func (l *dbList) remove (e *Node) *Node {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil//避免内存泄露
	e.prev = nil//避免内存泄露
	e.list = nil
	l.len --
	return  e
}
//头插
func (l *dbList) PushFront(v interface{}) *Node {
	l.lazyInit()
	node := &Node{Value: v}
	return l.insert(node, &l.head)
}
//尾插
func (l *dbList) PushBack(v interface{}) *Node {
	l.lazyInit()
	node := &Node{Value: v}
	return l.insert(node, l.head.prev)
}

//打印链表
func (l *dbList)print() {
	if l.len == 0 {
		return
	}
	cur := l.head.next
	for cur !=  &l.head {
		fmt.Println(cur.Value.(string))
		cur = cur.next
	}
}


func main() {
  dbList := New()
  dbList.PushBack("a")
  dbList.PushBack("b")
  dbList.PushBack("ac")
  dbList.print()
}