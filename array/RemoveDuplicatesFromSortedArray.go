package main

import "fmt"

/*
----------------------------------------------------------------------------
给定一个排序数组 你需要在原地删除重复的元素 使得每个元素出现一次 返回移除后的新长度
不要使用额外的空间 保证空间复杂度为O(1)
---------------------------------------------------------------------------
如给定数组 {1,2,2,3}
应该返回{1,2,3}新数组的长度 即3
-----------------------------
 */

func main()  {
	var arr = []int{1,1,2,2,3,4,4}
	re := removeDuplicatesFromSortedArray(arr)
	fmt.Println(re)
}

func removeDuplicatesFromSortedArray(arr []int)  int {
	// 如果为空
	if len(arr) <= 1 {
		return len(arr)
	}

	// 定义最后不重复元素指针
	slow := 0

	// 循环数组
	for i := 0;i < len(arr); i ++ {
		// 如果当前快指针节点不等于慢指针的值 就将慢指针的值+1 并且用快指针的元素值赋值给慢指针的位置 如果等于,拿慢指针没必要移动了 继续保持那个位置
		if arr[slow] != arr[i] {
			slow++
			arr[slow] = arr[i]
		}
	}

	return slow + 1
}
