package main

import (
	"dataStructure/linkList"
	"fmt"
)

func main() {

		var nodeOne = &linkList.ListNode{}
	    var nodeTwo = &linkList.ListNode{}

	nodeOne.TailAdd(2)
	nodeOne.TailAdd(4)
	nodeOne.TailAdd(6)

	nodeTwo.TailAdd(3)
	nodeTwo.TailAdd(5)
	nodeTwo.TailAdd(6)

	// 合并？？
	currentOne := nodeOne
	currentTwo := nodeTwo

	fmt.Println("合并前第一个链表:")
	linkList.PrintNode(nodeOne)
	fmt.Println("合并前第二个链表:")
	linkList.PrintNode(nodeTwo)

	var mergeList = &linkList.ListNode{}
	for {
		if currentTwo.Next == nil  || currentOne.Val <= currentTwo.Val{ //第二个链表无值
			mergeList.TailAdd(currentOne.Val)
			currentOne = currentOne.Next
		} else {
			mergeList.TailAdd(currentTwo.Val)
			currentTwo = currentTwo.Next
		}

		if currentOne.Next == nil && currentTwo.Next == nil {
			break
		}
	}

	fmt.Println("合并后链表:")
	linkList.PrintNode(mergeList)
}


