package linear

//单链表节点
type LinkedNode struct {
	//数据
	Data interface{}
	//指向下一个节点
	Next *LinkedNode
}

//单链表
type SingleLinkedList struct {
	Size   uint
	Header *LinkedNode
	Tail   *LinkedNode
}

//初始化单链表
func NewSingleLinkedList() (link *SingleLinkedList) {
	link = &SingleLinkedList{
		Size:   0,
		Header: nil,
		Tail:   nil,
	}
	return
}

//单链表尾部插入
func (s *SingleLinkedList) Append(data *LinkedNode) (ok bool) {
	if data == nil {
		return false
	}

	data.Next = nil

	if s.Size == 0 {
		s.Header = data
	} else {
		oldTail := s.Tail
		oldTail.Next = data
	}

	s.Tail = data
	s.Size++
	return true
}

//指定 index 插入节点数据
func (s *SingleLinkedList) Insert(index uint, data *LinkedNode) (ok bool) {

	if s.Size < index || s.Size == 0 {
		return false
	}

	//如果是最前插入
	if index == 0 {
		data.Next = s.Header
		s.Header = data
	} else {
		//找到插入位置的前一个节点
		preNode := s.Header
		for i := 1; i < int(index); i++ {
			preNode = preNode.Next
		}
		//插入数据
		data.Next = preNode.Next
		preNode.Next = data

		//如果是最后插入
		if index == s.Size-1 {
			s.Tail = preNode.Next
		}
	}

	s.Size++

	return true
}

//通过 index 删除链表数据
func (s *SingleLinkedList) Delete(index uint) (ok bool) {

	if s.Size < index || s.Size == 0 {
		return false
	}

	if index == 0 {
		nextNode := s.Header.Next
		s.Header = nextNode
		if index == s.Size-1 {
			s.Tail = nil
		}
	} else {
		//找到删除节点的前一个节点
		preNode := s.Header
		for i := 1; i < int(index); i++ {
			preNode = preNode.Next
		}
		//获取当前节点
		curNode := preNode.Next

		//判断是不是链表的最后一个数据
		if curNode.Next != nil {
			//获取当前节点下一个节点
			nextNode := curNode.Next
			//修改删除节点前一个节点的下一个节点
			preNode.Next = nextNode
		} else {
			preNode.Next = nil
		}

		if index == s.Size-1 {
			s.Tail = preNode
		}
	}

	s.Size--

	return true
}

//通过 index 获取链表数据
func (s *SingleLinkedList) Get(index uint) (data *LinkedNode) {

	//判断下标是否越界
	if index >= s.Size {
		return nil
	}

	//遍历查询对应的节点
	node := s.Header
	for i := 0; i < int(index); i++ {
		node = node.Next
	}

	return node
}

//获取链表 size
func (s *SingleLinkedList) GetSize() (size uint) {
	return s.Size
}

//移除链表中的所有节点
func (s *SingleLinkedList) RemoveAll() (ok bool) {
	s.Header = nil
	s.Tail = nil
	s.Size = 0
	return
}

//获取链表的头
func (s *SingleLinkedList) GetHeader() (node *LinkedNode) {
	return s.Header
}

//获取链表的尾
func (s *SingleLinkedList) GetTail() (node *LinkedNode) {
	return s.Tail
}
