package gotype

//用于演示Golang相关算法所写的数据结构
import (
	"sync"
)

type SliceQueue struct {
	arr []interface{}
	sync.RWMutex
}

func NewSliceQueue() *SliceQueue {
	return &SliceQueue{arr: make([]interface{}, 0)}
}

// 判断队列是否为空
func (p *SliceQueue) IsEmpty() bool {
	return p.Size() == 0
}

//返回队列的大小
func (p *SliceQueue) Size() int {
	return len(p.arr)
}

//返回队列首元素
func (p *SliceQueue) GetFront() interface{} {
	if p.IsEmpty() {
		return nil
	}
	return p.arr[0]
}

//返回队列尾元素
func (p *SliceQueue) GetBack() interface{} {
	if p.IsEmpty() {
		return nil
	}
	return p.arr[p.Size()-1]
}

//返回并移除队列尾元素
func (p *SliceQueue) PopBack() interface{} {
	p.Lock()
	defer p.Unlock()
	if p.IsEmpty() {
		return nil
	}
	ret := p.arr[p.Size()-1]
	p.arr = p.arr[:p.Size()-1]
	return ret
}

//删除队列头元素
func (p *SliceQueue) DeQueue() interface{} {
	p.Lock()
	defer p.Unlock()
	if len(p.arr) != 0 {
		first := p.arr[0]
		p.arr = p.arr[1:]
		return first
	} else {
		return nil
	}
}

//把新元素加入队列尾
func (p *SliceQueue) EnQueue(item interface{}) {
	p.Lock()
	defer p.Unlock()
	p.arr = append(p.arr, item)
}

//把新元素加入队列首
func (p *SliceQueue) EnQueueFirst(item interface{}) {
	p.Lock()
	defer p.Unlock()
	newQueue := []interface{}{item}
	p.arr = append(newQueue, p.arr[:]...)
}

//简单实现一个Remove
func (p *SliceQueue) Remove(item interface{}) {
	p.Lock()
	defer p.Unlock()
	for k, v := range p.arr {
		if v == item {
			p.arr = append(p.arr[:k], p.arr[k+1:]...)
		}
	}
}
func (p *SliceQueue) List() []interface{} {
	return p.arr
}

type LinkedQueue struct {
	head *LNode
	end  *LNode
	sync.RWMutex
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{}
}

//判断队列是否为空,如果为空返回true，否则返回false
func (p *LinkedQueue) IsEmpty() bool {
	return p.head == nil
}

//获取栈中元素的个数
func (p *LinkedQueue) Size() int {
	size := 0
	node := p.head
	for node != nil {
		node = node.Next
		size++
	}
	return size
}

//入队列：把元素e加到队列尾
func (p *LinkedQueue) EnQueue(e interface{}) {
	p.Lock()
	defer p.Unlock()
	node := &LNode{Data: e}
	if p.head == nil {
		p.head = node
		p.end = node
	} else {
		p.end.Next = node
		p.end = node
	}
}

//出队列，删除队列首元素
func (p *LinkedQueue) DeQueue() interface{} {
	p.Lock()
	defer p.Unlock()
	if p.head == nil {
		return nil
	}
	res := p.head
	p.head = p.head.Next
	if p.head == nil {
		p.end = nil
	}
	return res
}

//取得队列首元素
func (p *LinkedQueue) GetFront() interface{} {
	if p.head == nil {
		return nil
	}
	return p.head.Data
}

//取得队列尾元素
func (p *LinkedQueue) GetBack() interface{} {
	if p.end == nil {
		return nil
	}
	return p.end.Data
}
