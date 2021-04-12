package main

import (
	"fmt"
	"strconv"
)

////////////////////////////
// 简单的不重复整形单链表结构体
///////////////////////////

//单链表的node节点
type ListNode struct {
	value int
	next *ListNode
}

// 头部添加
func (node *ListNode) headAdd (i int) *ListNode {
	headNode := &ListNode{
		i,
		node,
	}

	return headNode
}

// 尾部添加
func (node *ListNode) tailAdd (i int) *ListNode {
	currentNode := node
	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	tailNode := currentNode

	newTailNode := &ListNode{
		i,
		nil,
	}

	tailNode.next = newTailNode
	return node
}

// 查找
func (node *ListNode) findVal (i int) bool  {
	currentNode := node
	for currentNode.next != nil{
		if currentNode.value == i {
			return true
			break
		}
		currentNode = currentNode.next
	}

	if currentNode.value == i {
		return true
	}

	return false
}

// 删除
func (node *ListNode) deleteVal (i int) * ListNode {
	currentNode := node

	for {
		if currentNode.value == i {
			return node.next
		}

		if currentNode.next.value == i {
			currentNode.next = currentNode.next.next
			return node
		}

		if currentNode.next == nil {
			break
		}

		currentNode = currentNode.next
	}

	return node
}

func main() {
	// 定义空链表
	node := &ListNode{
		0,
		nil,
	}

	// 尾部添加10个元素
	for i := 1; i <= 10; i++ {
		node.tailAdd(i)
	}
	printNode(node)

	// 头部添加3个元素
	for i := -3; i < 0; i++ {
		node = node.headAdd(i)
	}
	printNode(node)


	// 拆分链表
	fmt.Println("翻转链表前链表---------:")
	printNode(node)




	var prev *ListNode
	current := node
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}



	fmt.Println("翻转链表前链后----------:")
	printNode(prev)

	fmt.Printf("递归翻转前-----: \n")
	printNode(prev)

	list := recursionRevertLinkList(prev)

	fmt.Println("递归翻转后")
	printNode(list)
}

// 递归翻转单链表
func recursionRevertLinkList(node *ListNode)  *ListNode {
	// 递归首先确定的是返回
	if node == nil || node.next == nil {
		return node
	}

	// 依次递归 实际上从链表的最后开始
	newHead := recursionRevertLinkList(node.next)
	node.next.next = node
	node.next = nil

	return  newHead
}

func printNode(node *ListNode) {
	str := "单链表元素为:"
	currentNode := node
	for {

		str = str + strconv.Itoa(currentNode.value) + " "

		if currentNode.next == nil {
			break
		}

		currentNode = currentNode.next
	}

	fmt.Println(str)

	return
}
