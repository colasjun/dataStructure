package linkList

import (
"fmt"
"strconv"
)

////////////////////////////
// 简单的不重复整形单链表结构体
///////////////////////////

//单链表的node节点
type ListNode struct {
	Val int
	Next *ListNode
}

// 头部添加
func (node *ListNode) HeadAdd (i int) *ListNode {
	headNode := &ListNode{
		i,
		node,
	}

	return headNode
}

// 尾部添加
func (node *ListNode) TailAdd (i int) *ListNode {
	if node == nil {
		return &ListNode{
			i,
			nil,
		}
	}

	currentNode := node



	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	tailNode := currentNode

	newTailNode := &ListNode{
		i,
		nil,
	}

	tailNode.Next = newTailNode
	return node
}

// 查找
func (node *ListNode) FindVal (i int) bool  {
	currentNode := node
	for currentNode.Next != nil{
		if currentNode.Val == i {
			return true
			break
		}
		currentNode = currentNode.Next
	}

	if currentNode.Val == i {
		return true
	}

	return false
}

// 删除
func (node *ListNode) DeleteVal (i int) * ListNode {
	currentNode := node

	for {
		if currentNode.Val == i {
			return node.Next
		}

		if currentNode.Next.Val == i {
			currentNode.Next = currentNode.Next.Next
			return node
		}

		if currentNode.Next == nil {
			break
		}

		currentNode = currentNode.Next
	}

	return node
}

func PrintNode(node *ListNode) {
	str := "单链表元素为:"
	currentNode := node
	for {

		str = str + strconv.Itoa(currentNode.Val) + " "

		if currentNode.Next == nil {
			break
		}

		currentNode = currentNode.Next
	}

	fmt.Println(str)

	return
}

