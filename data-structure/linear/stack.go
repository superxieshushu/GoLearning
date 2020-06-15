package linear

//栈的实现：后进先出、使用单向列表实现、头插入即可

type Stack struct {
	list *SingleLinkedList
}

func NewStack() (s *Stack) {
	return &Stack{list: NewSingleLinkedList()}
}

//出栈
func (s *Stack) Pop() (data interface{}) {
	if s.list.Size == 0 {
		return
	}
	data = s.list.getTail().Data
	s.list.Delete(s.list.Size - 1)
	return
}

//压栈
func (s *Stack) Push(data interface{}) (ok bool) {
	s.list.Append(&LinkedNode{Data: data})
	return
}

//获取栈顶的数据
func (s *Stack) Peek() (data interface{}) {
	if s.list.Tail != nil {
		data = s.list.Tail.Data
	}
	return
}

//清空栈
func (s *Stack) Clear() {
	s.list.RemoveAll()
}

//获取栈的长度
func (s *Stack) Length() (size uint) {
	return s.list.Size
}
