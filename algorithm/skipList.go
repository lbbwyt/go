// skipList
// 跳表的实现
package algorithm

import (
	"fmt"
	"math/rand"
)

//跳表时基于有序链表实现的，
type Node struct {
	Forward []Node //Forward数组的大学即该节点的层数
	Value   interface{}
}

const SKIPLIST_MAXLEVEL = 32 //跳表的最大层级
const SKIPLIST_P = 4

func NewNode(v interface{}, level int) *Node {
	return &Node{Value: v, Forward: make([]Node, level)}
}

type SkipList struct {
	Header *Node //头结点
	Level  int   //跳表对应的层级
}

func NewSkipList() *SkipList {
	return &SkipList{
		Header: NewNode(0, SKIPLIST_MAXLEVEL),
		Level:  1,
	}
}

func (skipList *SkipList) Insert(key int) {
	//update 暂存临时节点
	update := make(map[int]*Node)
	//node为当前节点
	node := skipList.Header
	//逐层遍历,找到每层要插入位置的前一个节点，并存在update中
	//	跳表的级数，简单来解释，就是层数-1（最下面一级是0级）。
	for i := skipList.Level - 1; i >= 0; i-- {
		for {
			//node.Forward[i]表示当前层（即第i层），该节点对应的下一个节点
			if node.Forward[i].Value != nil && node.Forward[i].Value.(int) < key {
				node = &node.Forward[i]
			} else {
				break
			}
		}
		update[i] = node
	}

	level := skipList.Random_level()
	if level > skipList.Level {
		for i := skipList.Level; i < level; i++ {
			update[i] = skipList.Header
		}
		skipList.Level = level
	}
	newNode := NewNode(key, level)

	for i := 0; i < level; i++ {
		newNode.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = *newNode
	}
}

//插入节点时随机生成该节点的层数
func (skipList *SkipList) Random_level() int {

	level := 1
	for (rand.Int31()&0xFFFF)%SKIPLIST_P == 0 {
		level += 1
	}
	if level < SKIPLIST_MAXLEVEL {
		return level
	} else {
		return SKIPLIST_MAXLEVEL
	}
}

//打印跳表
func (skipList *SkipList) PrintSkipList() {

	fmt.Println("\nSkipList-------------------------------------------")
	for i := SKIPLIST_MAXLEVEL - 1; i >= 0; i-- {

		fmt.Println("level:", i)
		node := skipList.Header.Forward[i]
		for {
			if node.Value != nil {
				fmt.Printf("%d ", node.Value.(int))
				node = node.Forward[i]
			} else {
				break
			}
		}
		fmt.Println("\n--------------------------------------------------------")
	}

	fmt.Println("Current MaxLevel:", skipList.Level)
}

func (skipList *SkipList) Search(key int) *Node {

	node := skipList.Header
	//从最顶层开始搜索
	for i := skipList.Level - 1; i >= 0; i-- {

		for {
			if node.Forward[i].Value == nil {
				break
			}

			if node.Forward[i].Value.(int) == key {
				return &node.Forward[i]
			}

			if node.Forward[i].Value.(int) < key {
				node = &node.Forward[i]
				continue
			} else { // > key
				break
			}
		}

	}
	return nil
}

//和插入类似，需先找到每层要删除节点的前一个节点
func (skipList *SkipList) Remove(key int) {

	update := make(map[int]*Node)
	node := skipList.Header
	for i := skipList.Level - 1; i >= 0; i-- {

		for {

			if node.Forward[i].Value == nil {
				break
			}

			if node.Forward[i].Value.(int) == key {
				update[i] = node
				break
			}

			if node.Forward[i].Value.(int) < key {
				node = &node.Forward[i]
				continue
			} else { // > key
				break
			}

		}

	}

	for i, v := range update {
		if v == skipList.Header {
			skipList.Level--
		}
		v.Forward[i] = v.Forward[i].Forward[i]
	}
}
