package gotype

import (
	"fmt"
	"github.com/isdamir/assert"
	"testing"
)

func TestSliceQueue(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Slice构建队列结构")
	queue := NewSliceQueue()
	queue.EnQueue(1)
	queue.EnQueue(2)
	assert.Equal(t, 1, queue.GetFront())
	assert.Equal(t, 2, queue.GetBack())
	assert.Equal(t, 2, queue.Size())
}
func TestLinkedQueue(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Slice构建队列结构")
	queue := NewLinkedQueue()
	queue.EnQueue(1)
	queue.EnQueue(2)
	assert.Equal(t, 1, queue.GetFront())
	assert.Equal(t, 2, queue.GetBack())
	assert.Equal(t, 2, queue.Size())
}
