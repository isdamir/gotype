package gotype

import (
	"fmt"
)
//定义用于演示算法时的一些公共方法

//创建链表
func CreateNode(node *LNode,max int) {
	cur := node
	for i := 1; i < max; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
}

//打印链表的方法
func PrintNode(info string, node *LNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}

func ArrayToTree(arr []int,start int,end int) *BNode{
	var root *BNode
	if end>=start{
		root=&BNode{}
		mid:=(start+end+1)/2
		//树的根结点为数组中间的元素
		root.Data=arr[mid]
		//递归的用左半部分数组构造root的左子树
		root.LeftChild=ArrayToTree(arr,start,mid-1)
		//递归的用右半部分数组构造root的右子树
		root.RightChild=ArrayToTree(arr,mid+1,end)
	}
	return root
}
//用中序遍历的方式打印出二叉树结点的内容
func PrintTreeMidOrder(root *BNode){
	if root==nil{
		return
	}
	//遍历root结点的左子树
	if root.LeftChild!=nil{
		PrintTreeMidOrder(root.LeftChild)
	}
	//遍历root结点
	fmt.Print(root.Data," ")
	//遍历root结点的右子树
	if root.RightChild!=nil{
		PrintTreeMidOrder(root.RightChild)
	}
}