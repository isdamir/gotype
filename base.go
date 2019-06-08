package gotype

//链表定义
type LNode struct {
	Data interface{}
	Next *LNode
}

func NewLNode() *LNode {
	return &LNode{}
}

//二叉树定义
type BNode struct {
	Data       interface{}
	LeftChild  *BNode
	RightChild *BNode
}

func NewBNode() *BNode {
	return &BNode{}
}

//Trie树
type TrieNode struct {
	IsLeaf bool
	Url    string
	Child  []*TrieNode
}

func NewTrieNode(count int) *TrieNode {
	return &TrieNode{IsLeaf: false, Url: "", Child: make([]*TrieNode, count)}
}

//AVL树
type AVLNode struct {
	Data   int
	Height int
	Count  int
	Left   *AVLNode
	Right  *AVLNode
}

func NewAVLNode(data int) *AVLNode {
	return &AVLNode{data, 1, 1, nil, nil}
}
