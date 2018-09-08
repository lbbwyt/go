package main

import (
	"fmt"
)

type QueueByStack struct {
	in  *Stack
	out *Stack
}

//插入队列
func (queue *QueueByStack) push(data interface{}) {
	queue.in.Push(data)
}

func (queue *QueueByStack) pop() interface{} {
	if queue.out.Size() == 0 {
		for queue.in.Size() > 0 {
			queue.out.Push(queue.in.Pop())
		}
	}

	if queue.out.Size() == 0 {
		panic("队列为空")
	}
	return queue.out.Pop()

}

func main() {
	in := NewStack()
	out := NewStack()
	queue := &QueueByStack{
		in:  in,
		out: out,
	}
	queue.push(3)
	queue.push(2)
	queue.push(1)

	fmt.Println(queue.pop().(int))
	fmt.Println(queue.pop().(int))
	fmt.Println(queue.pop().(int))

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
