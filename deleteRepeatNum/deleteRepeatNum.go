package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	// 双指针算法确定数组不重复元素个数 已经排好序的数组
	arr := []int{1,2,2,3,3,4,4}
	num := deleteRepeatNum(arr)
	fmt.Printf("%v 中不重复元素为%d个 \n", arr, num)
}

func deleteRepeatNum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	i := 0
	for j := 1;j <len(arr);j++ {
		if arr[i] != arr[j] {
			i++
			arr[i] = arr[j]
		}
	}

	return i + 1
}


func getArr(num int) []int {
	var arr []int
	var i = 0
	rand.Seed(time.Now().UnixNano())
	for  i < num {
		arr = append(arr, rand.Intn(100))
		i ++
	}
	return arr
}
