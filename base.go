package gotype
//链表定义
type LNode struct {
	Data interface{}
	Next *LNode
}
func NewLNode()*LNode{
	return &LNode{}
}
//二叉树定义
type BNode struct{
	Data interface{}
	LeftChild *BNode
	RightChild *BNode
}
func NewBNode()*BNode{
	return &BNode{}
}