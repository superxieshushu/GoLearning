package linear

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {

	stack := NewStack()

	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	fmt.Println(stack.Pop())

	fmt.Println(stack.Peek())

	fmt.Println(stack.Length())
}
