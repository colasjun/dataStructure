package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	hr()
	arr := getArr(25)
	fmt.Println("冒泡排序前", arr)
	sort := bubbleSort(arr, len(arr))
	verifySort(arr, sort)
	fmt.Println("冒泡排序后", sort)
	hr()
	arr = getArr(25)
	fmt.Println("选择排序前", arr)
	sort = selectSort(arr, len(arr))
	verifySort(arr, sort)
	fmt.Println("选择排序后", sort)
	hr()
	arr = getArr(25)
	fmt.Println("插入排序前", arr)
	sort = insertSort(arr, len(arr))
	verifySort(arr, sort)
	fmt.Println("插入排序后", sort)
	hr()
	arr = getArr(25)
	fmt.Println("希尔排序前", arr)
	sort = shellSort(arr, len(arr))
	verifySort(arr, sort)
	fmt.Println("希尔排序后", sort)

}

func shellSort(arr []int, len int) []int {
	// 先算出步长数组
	step := len
	if (step >> 1) > 0 {
		step = step >> 1
	}

	fmt.Println(step)
	for step > 0 {
		// 快排
		for i := 0; i < len; i += step {
			//cur := i
			//while cur
			if arr[i] > arr[i+step] {
				arr[i], arr[i+step] = arr[i+step],arr[i]
			}
		}

		step >>= 1
	}


	return arr
}



func insertSort(arr []int, len int) []int {
	// 从第一值开始 后面每遍历一个数字 就将数据和前面已经排好序的比较 插入合适的位置即可
	for i := range arr {
		index := i - 1
		curVal := arr[i]
		for index >= 0 && curVal < arr[index]  {
			arr[index+1] = arr[index]
			index--
		}
		arr[index+1] = curVal
	}
	return arr
}


func selectSort(arr []int, len int) []int {
	// 遍历一遍所有数组 拿出最值交换到尾部 再继续遍历
	for i := 1; i <= len ; i ++ {
		maxKey := 0
		for j := 1; j <= len - i; j++ {
			if arr[j] > arr[maxKey] {
				maxKey = j
			}
		}


		arr[maxKey],arr[len-i] = arr[len-i],arr[maxKey]
	}

	return arr
}

func bubbleSort (arr []int, len int) []int {
	// 位置依次比较
	for i := 1; i <= len; i ++ {
		for j := 1; j <= len - i; j ++ {
			if arr[j-1] > arr[j]  {
				arr[j-1], arr[j] = arr[j],arr[j-1]
			}
		}
	}
	return arr
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

// 验证排序是否正确
func verifySort(arr []int,sort []int) {
	isOk := true

	for i := 0; i < len(sort) - 1; i++ {
		if sort[i] > sort[i+1] {
			isOk = false
		}
	}

	if !isOk {
		fmt.Println("排序异常")
		return
	}


	// 检测数组之和是否ok
	sumArr := 0
	sumSort := 0
	for _,k := range arr{
		sumArr+=k
	}
	for _,k := range sort{
		sumSort+=k
	}

	if sumArr == sumSort {
		fmt.Println("ok")
		return
	}

	fmt.Println("排序异常")
	return
}

func hr()  {
	fmt.Println()
	fmt.Println("------------------------------------------------------------------------------------------------")
}