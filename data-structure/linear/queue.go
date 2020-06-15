package linear

//队列实现：先进先出
type Queue struct {
	list *SingleLinkedList
}

//新建初始化队列
func NewQueue() (q *Queue) {
	return &Queue{list: NewSingleLinkedList()}
}

//对尾插入数据
func (q *Queue) Put(data interface{}) {
	q.list.Append(&LinkedNode{Data: data})
}

//队首取出数据
func (q *Queue) Poll() (data interface{}) {
	if q.list.Header == nil {
		return
	}
	data = q.list.Header.Data
	q.list.Delete(0)
	return
}

//清除队列
func (q *Queue) Clear() {
	q.list.RemoveAll()
}
