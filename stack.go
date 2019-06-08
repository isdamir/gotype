package gotype

//用于演示Golang相关算法所写的数据结构
//也可以使用container/list 来实现，本质是双向链表
import (
	"sync"
)

//slice所写的stack
type SliceStack struct {
	arr       []interface{}
	stackSize int
	sync.RWMutex
}

func NewSliceStack() *SliceStack {
	return &SliceStack{arr: make([]interface{}, 0)}
}

//判断是否为空
func (p *SliceStack) IsEmpty() bool {
	return p.stackSize == 0
}

//返回大小
func (p *SliceStack) Size() int {
	return p.stackSize
}

//返回栈顶元素
func (p *SliceStack) Top() interface{} {
	if p.IsEmpty() {
		return nil
	}
	return p.arr[p.stackSize-1]
}

//弹出栈元素
func (p *SliceStack) Pop() interface{} {
	p.Lock()
	defer p.Unlock()
	if p.stackSize > 0 {
		p.stackSize--
		ret := p.arr[p.stackSize]
		p.arr = p.arr[:p.stackSize]
		return ret
	}
	return nil
}

//Push栈元素
func (p *SliceStack) Push(t interface{}) {
	p.Lock()
	defer p.Unlock()
	p.arr = append(p.arr, t)
	p.stackSize = p.stackSize + 1
}
func (p *SliceStack) List() []interface{} {
	return p.arr
}

//链表所写的stack
type LinkedStack struct {
	head *LNode
	sync.RWMutex
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{head: &LNode{}}
}
func (p *LinkedStack) IsEmpty() bool {
	return p.head.Next == nil
}
func (p *LinkedStack) Size() int {
	size := 0
	node := p.head.Next
	for node != nil {
		node = node.Next
		size++
	}
	return size
}
func (p *LinkedStack) Push(e interface{}) {
	p.Lock()
	defer p.Unlock()
	node := &LNode{Data: e, Next: p.head.Next}
	p.head.Next = node
}
func (p *LinkedStack) Pop() interface{} {
	p.Lock()
	defer p.Unlock()
	tmp := p.head.Next
	if tmp != nil {
		p.head.Next = tmp.Next
		return tmp.Data
	}
	return nil
}
func (p *LinkedStack) Top() interface{} {
	if p.head.Next != nil {
		return p.head.Next.Data
	}
	return nil
}
