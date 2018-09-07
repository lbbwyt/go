package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(data int) *TreeNode {
	return &TreeNode{
		value: data,
		left:  nil,
		right: nil,
	}
}

//求树的深度
//树是一种递归结构，很多树的问题可以使用递归来处理。
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(maxDepth(root.left), maxDepth(root.right)) + 1
}

//反转树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := root.left //// 后面的操作会改变 left 指针，因此先保存下来
	root.left = invertTree(root.right)
	root.right = invertTree(left)
	return root

}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmet(root.left, root.right)
}

func isSymmet(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.value != t2.value {
		return false
	}
	return isSymmet(t1.left, t2.right) && isSymmet(t1.right, t2.left)
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//树的层次遍历
//前序遍历为 root -> left -> right，后序遍历为 left -> right -> root。可以修改前序遍历成为 root -> right -> left，那么这个顺序就和后序遍历正好相反。
func levelTraversal(root *TreeNode) {
	queue := list.New()
	queue.PushBack(root)
	for queue.Front() != nil {
		t := queue.Front()
		queue.Remove(t)
		temp := t.Value.(*TreeNode)
		fmt.Println(temp.value)
		if temp.left != nil {
			queue.PushBack(temp.left)
		}
		if temp.right != nil {
			queue.PushBack(temp.right)
		}
	}
}

//前序遍历为 root -> left -> right，

func preTraversal(root *TreeNode) {
	stack := NewStack()
	stack.Push(root)
	for stack.Size() > 0 {
		node := stack.Pop().(*TreeNode)
		if node == nil {
			continue
		}
		fmt.Println(node.value)
		stack.Push(node.right) // 先右后左，保证左子树先遍历
		stack.Push(node.left)
	}
}

//后序遍历为 left -> right -> root。
//可以修改前序遍历成为 root -> right -> left，那么这个顺序就和后序遍历正好相反。

func postTraversal(root *TreeNode) {
	stack := NewStack()
	stack2 := NewStack()
	stack.Push(root)
	for stack.Size() > 0 {
		node := stack.Pop().(*TreeNode)
		if node == nil {
			continue
		}
		stack2.Push(node.value)
		stack.Push(node.left)
		stack.Push(node.right)
	}
	stack3 := stack2
	for stack2.Size() > 0 {
		fmt.Println(stack3.Pop())
	}
}

//中遍历为 left -> root -> right，

func inorderTraversal(root *TreeNode) {
	stack := NewStack()
	cur := root
	for cur != nil || stack.Size() > 0 {
		for cur != nil {
			stack.Push(cur)
			cur = cur.left
		}

		node := stack.Pop().(*TreeNode)
		fmt.Println(node.value)
		cur = node.right
	}
}

func main() {
	tree1 := NewTreeNode(2)
	tree1.right = NewTreeNode(4)
	tree1.left = NewTreeNode(8)
	t2 := tree1.left
	t2.right = NewTreeNode(5)
	fmt.Println(isSymmetric(tree1))
	//	levelTraversal(tree1)
	inorderTraversal(tree1)
}

//栈及其操作
type Stack struct {
	Element []interface{} //Element
}

func NewStack() *Stack {
	return &Stack{}
}

//压入栈，支持压入多个
func (stack *Stack) Push(value ...interface{}) {
	stack.Element = append(stack.Element, value...)
}

//返回下一个元素,并从Stack移除元素
func (stack *Stack) Pop() interface{} {
	if stack.Size() > 0 {
		t := stack.Element[stack.Size()-1]
		stack.Element = stack.Element[:stack.Size()-1]
		return t
	}
	return nil //read empty stack
}

func (stack *Stack) Size() int {
	return len(stack.Element)
}
