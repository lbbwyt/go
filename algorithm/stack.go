package algorithm

import (
	"errors"
)

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

//返回下一个元素
func (stack *Stack) Top() (value interface{}) {
	if stack.Size() > 0 {
		return stack.Element[stack.Size()-1]
	}
	return nil //read empty stack
}

//Stack的size
func (stack *Stack) Size() int {
	return len(stack.Element)
}

//返回下一个元素,并从Stack移除元素
func (stack *Stack) Pop() (err error) {
	if stack.Size() > 0 {
		stack.Element = stack.Element[:stack.Size()-1]
		return nil
	}
	return errors.New("Stack为空.") //read empty stack
}
