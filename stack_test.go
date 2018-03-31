package gotype

import (
	"testing"
	"fmt"
	"github.com/isdamir/assert"
)
//测试Slice的Stack
func TestSliceStatck(t *testing.T){
	defer func(){
		if err:=recover();err!=nil{
            fmt.Println(err)
        }
	}()
	stack:=NewSliceStack()
	stack.Push(1)
	assert.Equal(t,1,stack.Top().(int))
	assert.Equal(t,1,stack.Size())
	stack.Pop()
	assert.Equal(t,0,stack.Size())
}
//测试Linked的Stack
func TestLinkedStatck(t *testing.T){
	defer func(){
		if err:=recover();err!=nil{
            fmt.Println(err)
        }
	}()
	stack:=NewLinkedStack()
	stack.Push(1)
	assert.Equal(t,1,stack.Top())
	assert.Equal(t,1,stack.Size())
	stack.Pop()
	assert.Equal(t,0,stack.Size())
}