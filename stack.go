package gotype
//用于演示Golang相关算法所写的数据结构
import (
	"sync"
	"errors"
)
//slice所写的stack
type SliceStack struct {
	arr []interface{}
	stackSize int
	sync.RWMutex
}
func NewSliceStack()*SliceStack{
	return &SliceStack{arr:make([]interface{},0)}
}
//判断是否为空
func (p *SliceStack) IsEmpty()bool{
	return p.stackSize==0
}
//返回大小
func (p *SliceStack)  Size() int{
	return p.stackSize
}
//返回栈顶元素
func (p *SliceStack) Top() interface{}{
	if(p.IsEmpty()){
		panic(errors.New("栈已经为空"))
	}
	return p.arr[p.stackSize-1]
}
//弹出栈元素
func (p *SliceStack) Pop() interface{}{
	p.Lock()
	defer p.Unlock()
	if(p.stackSize>0){
		p.stackSize--
		ret:=p.arr[p.stackSize]
		p.arr=p.arr[:p.stackSize]
		return ret
	}
	panic(errors.New("栈已经为空"))
}
//Push栈元素
func (p *SliceStack) Push(t interface{}){
	p.Lock()
	defer p.Unlock()
	p.arr=append(p.arr,t)
	p.stackSize =p.stackSize+1
}

//链表所写的stack
type LinkedStack struct{
	head *LNode
	sync.RWMutex
}
func NewLinkedStack()*LinkedStack{
	return &LinkedStack{head:&LNode{}}
}
func (p *LinkedStack) IsEmpty() bool{
	return p.head.next==nil
}
func (p *LinkedStack) Size() int{
	size:=0
	node:=p.head.next
	for node!=nil {
		node=node.next
		size++
	}
	return size
}
func (p *LinkedStack) Push(e interface{}){
	p.Lock()
	defer p.Unlock()
	node:=&LNode{data:e,next:p.head.next}
	p.head.next=node
}
func (p *LinkedStack) Pop() interface{}{
	p.Lock()
	defer p.Unlock()
	tmp:=p.head.next
	if tmp!=nil{
		p.head.next = tmp.next
		return tmp.data
	}
	panic(errors.New("栈已经为空"))
}
func (p *LinkedStack) Top() interface{}{
	if p.head.next!=nil{
		return p.head.next.data
	}
	panic(errors.New("栈已经为空"))
}