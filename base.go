package gotype
//链表定义
type LNode struct {
	data interface{}
	next *LNode
}
func NewLNode()*LNode{
	return &LNode{}
}
//二叉树定义
type BNode struct{
	data interface{}
	leftChild *BNode
	rightChild *BNode
}
func NewBNode()*BNode{
	return &BNode{}
}