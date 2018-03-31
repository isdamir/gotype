package gotype
//用于演示Golang相关算法所写的数据结构
import (
	"sync"
	"errors"
)
type SliceQueue struct{
	arr []interface{}
	front int //队列头
	rear int //队列尾
	sync.RWMutex
}
func NewSliceQueue()*SliceQueue{
	return &SliceQueue{arr:make([]interface{},0)}
}
// 判断队列是否为空
func (p *SliceQueue) IsEmpty() bool{
	return p.front==p.rear
}
//返回队列的大小
func (p *SliceQueue) Size() int{
	return p.rear-p.front
}
//返回队列首元素
func (p *SliceQueue) GetFront() interface{}{
	if p.IsEmpty(){
		panic(errors.New("队列已经为空"))
	}
	return p.arr[p.front]
}
//返回队列尾元素
func (p *SliceQueue) GetBack()interface{}{
	if p.IsEmpty(){
		panic(errors.New("队列已经为空"))
	}
	return p.arr[p.rear-1]
} 
//删除队列头元素
func (p *SliceQueue) DeQueue(){
	if p.rear>p.front{
		p.rear--
		p.arr=p.arr[1:]
	}else{
		panic(errors.New("队列已经为空"))
	}
}
//把新元素加入队列尾
func (p *SliceQueue) EnQueue(item interface{}){
	p.arr=append(p.arr,item)
	p.rear++
}
//简单实现一个Remove
func (p *SliceQueue) Remove(item interface{}) {
	for k, v := range p.arr {
		if v == item {
			p.arr = append(p.arr[:k], p.arr[k+1:]...)
			p.rear--
		}
	}
}
func (p *SliceQueue) List() []interface{} {
	return p.arr
}
type LinkedQueue struct{
	head *LNode
	end *LNode
}
func NewLinkedQueue()*LinkedQueue{
	return &LinkedQueue{}
}
//判断队列是否为空,如果为空返回true，否则返回false
func (p *LinkedQueue) IsEmpty() bool{
	return p.head==nil
}
//获取栈中元素的个数
func (p *LinkedQueue) Size() int{
	size:=0
	node:=p.head
	for node!=nil{
		node=node.next
		size++
	}
	return size
}
//入队列：把元素e加到队列尾
func (p *LinkedQueue) EnQueue(e interface{}){
	node:=&LNode{data:e}
	if p.head==nil{
		p.head=node
		p.end=node
	}else{
		p.end.next=node
		p.end=node
	}
}
//出队列，删除队列首元素
func (p *LinkedQueue) DeQueue(){
	if p.head==nil{
		panic(errors.New("队列已经为空"))
	}
	p.head=p.head.next
	if p.head==nil{
		p.end=nil
	}
}
//取得队列首元素
func (p *LinkedQueue) GetFront() interface{}{
	if p.head==nil{
		panic(errors.New("队列已经为空"))
	}
	return p.head.data
}
//取得队列尾元素
func (p *LinkedQueue) GetBack() interface{}{
	if p.end==nil{
		panic(errors.New("队列已经为空"))
	}
	return p.end.data
}