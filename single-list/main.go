package main

import (
	"fmt"
	"strconv"
)

////////////////////////////
// 简单的不重复整形单链表结构体
///////////////////////////

//单链表的node节点
type listNode struct {
	value int
	next *listNode
}

// 头部添加
func (node *listNode) headAdd (i int) *listNode {
	headNode := &listNode{
		i,
		node,
	}

	return headNode
}

// 尾部添加
func (node *listNode) tailAdd (i int) *listNode {
	currentNode := node
	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	tailNode := currentNode

	newTailNode := &listNode{
		i,
		nil,
	}

	tailNode.next = newTailNode
	return node
}

// 查找
func (node *listNode) findVal (i int) bool  {
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
func (node *listNode) deleteVal (i int) * listNode {
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
	node := &listNode{
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

	// 查找元素10
	isExist := node.findVal(10)
	fmt.Println("查找元素10结果:",isExist)


	// 查找元素11
	isExist = node.findVal(11)
	fmt.Println("查找元素11结果:",isExist)

	// 删除元素-1
	node = node.deleteVal(-1)
	printNode(node)

	// 删除元素10
	node = node.deleteVal(10)
	printNode(node)

	// 删除元素7
	node = node.deleteVal(7)
	printNode(node)
}

func printNode(node *listNode) {
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
