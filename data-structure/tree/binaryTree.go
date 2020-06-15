package tree

//定义树节点
type BinaryNode struct {
	Data   int64
	Left   *BinaryNode
	Right  *BinaryNode
	Parent *BinaryNode
}

//定义二叉树
type BinaryTree struct {
	size uint64
	root *BinaryNode
}

//初始化一个二叉数
func NewBinaryTree() (tree *BinaryTree) {
	return &BinaryTree{
		size: 0,
		root: nil,
	}
}

//插入一个树节点
func (b *BinaryTree) Append(node *BinaryNode) {

}

//遍历二叉树
func Traversal(tree *BinaryNode) {
	if tree == nil {
		return
	}
	Traversal(tree.Left)
	Traversal(tree.Right)
	print(tree.Data)

}

//二叉树的遍历
func main() {
	Traversal(nil)
}
