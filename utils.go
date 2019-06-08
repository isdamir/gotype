package gotype

import (
	"fmt"
)

//定义用于演示算法时的一些公共方法

//创建链表
func CreateNode(node *LNode, max int) {
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

//通过数组创建中序的二叉树
func ArrayToTree(arr []int, start int, end int) *BNode {
	var root *BNode
	if end >= start {
		root = &BNode{}
		mid := (start + end + 1) / 2
		//树的根结点为数组中间的元素
		root.Data = arr[mid]
		//递归的用左半部分数组构造root的左子树
		root.LeftChild = ArrayToTree(arr, start, mid-1)
		//递归的用右半部分数组构造root的右子树
		root.RightChild = ArrayToTree(arr, mid+1, end)
	}
	return root
}

//用中序遍历的方式打印出二叉树结点的内容
func PrintTreeMidOrder(root *BNode) {
	if root == nil {
		return
	}
	//遍历root结点的左子树
	if root.LeftChild != nil {
		PrintTreeMidOrder(root.LeftChild)
	}
	//遍历root结点
	fmt.Print(root.Data, " ")
	//遍历root结点的右子树
	if root.RightChild != nil {
		PrintTreeMidOrder(root.RightChild)
	}
}

//层序打印二叉树
func PrintTreeLayer(node *BNode) {
	if node == nil {
		return
	}
	var p *BNode
	queue := NewSliceQueue()
	//树根结点进队列
	queue.EnQueue(node)
	for queue.Size() > 0 {
		p = queue.DeQueue().(*BNode)
		//访问当前结点
		fmt.Print(p.Data, " ")
		//如果这个结点的左孩子不为空则入队列
		if p.LeftChild != nil {
			queue.EnQueue(p.LeftChild)
		}
		//如果这个结点的右孩子不为空则入队列
		if p.RightChild != nil {
			queue.EnQueue(p.RightChild)
		}
	}
}
func Abs(x int) int {
	switch {
	case x < 0:
		return -x
	case x == 0:
		return 0
	}
	return x
}
func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func Min3(a, b, c int) int {
	min := Min(a, b)
	return Min(min, c)
}
func Max3(a, b, c int) int {
	max := Max(a, b)
	return Max(max, c)
}

//判断数字n的二进制数从右往左数第i位是否为1
func IsOne(n, i int) bool {
	return (uint(n) & (uint(1) << uint(i))) == 1
}
func SwapInt(data []int, x, y int) {
	tmp := data[x]
	data[x] = data[y]
	data[y] = tmp
}
func SwapRune(data []rune, x, y int) {
	tmp := data[x]
	data[x] = data[y]
	data[y] = tmp
}
func IsNumber(c rune) bool {
	return c >= '0' && c <= '9'
}
