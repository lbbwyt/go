// trie
//字典树的实现
package logic

import "fmt"

//定义节点
type Node struct {
	char   rune           //字符，rune类型
	childs map[rune]*Node //孩子节点
	Data   interface{}    //该节点标识的数据
	deep   int            //节点深度
	isTerm bool           //并非叶子节点,仅标识该节点处是否存在字符串
}

//定义树
type Trie struct {
	root *Node //根节点
	size int
}

//新建一个节点

func NewNode(char rune, deep int) *Node {
	node := &Node{
		char:   char,
		deep:   deep,
		childs: make(map[rune]*Node, 16),
	}
	return node
}

//新建一个树
func NewTrie() *Trie {
	root := NewNode(' ', 1)
	return &Trie{
		root: root,
		size: 1,
	}
}

//添加一个节点
func (t *Trie) Add(key string, data interface{}) {
	curNode := t.root
	allChars := []rune(key)
	for _, char := range allChars {
		node, ok := curNode.childs[char] //判断当前节点的子节点存在该字符
		if !ok {
			//如果不存在，新建一个节点作为当前节点的子节点
			node = NewNode(char, curNode.deep+1)
			curNode.childs[char] = node
		}
		// 存在的话将子节点作为当前节点
		curNode = node
	}
	//	if len(curNode.childs < 1) {
	//	   curNode.isTerm = true
	//	}
	curNode.isTerm = true //给该节点打标记，
	curNode.Data = data
}

//前缀搜索

func (t *Trie) PrefixSearch(key string, limit int) (nodes []*Node) {
	curNode := t.root
	allChars := []rune(key)
	for _, char := range allChars {
		child, ok := curNode.childs[char]
		if !ok {
			fmt.Printf("prefix char:%c\n", char)
			return
		}

		curNode = child
	}

	//遍历以当前节点为根的子树，将isTrem为true的节点存在nodes中，
	//树的层次遍历，广度优先
	var queue []*Node
	queue = append(queue, curNode)

	for len(queue) > 0 {
		var q2 []*Node
		for _, n := range queue {
			if n.isTerm == true {
				nodes = append(nodes, n)
				if len(nodes) > limit {
					return
				}
			}
			for _, v := range n.childs {
				q2 = append(q2, v)
			}
		}
		queue = q2
	}
	return
}
