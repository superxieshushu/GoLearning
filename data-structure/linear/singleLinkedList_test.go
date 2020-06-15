package linear

import (
	"fmt"
	"testing"
)

func TestSingleLinkedList(t *testing.T) {

	linkedList := NewSingleLinkedList()

	linkedList.Append(&LinkedNode{Data: 1})

	linkedList.Append(&LinkedNode{Data: 2})

	ok := linkedList.Insert(1, &LinkedNode{Data: -2})
	if ok {
		fmt.Println("插入成功")
	}

	linkedList.Delete(2)

	fmt.Println(linkedList.Get(0))

	linkedList.RemoveAll()

	node := linkedList.Header

	for {
		if node != nil {
			fmt.Println(node.Data)
			node = node.Next
		} else {
			break
		}
	}
}
