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
	hr()
	arr = getArr(25)
	fmt.Println("堆排序前", arr)
	sort = heapSort(arr, len(arr))
	verifySort(arr, sort)
	fmt.Println("堆排序后", sort)
	//printTree(sort, len(sort), 0, 0)
}

/*func mergeSort(arr []int, len int)  {

}

func subMergeSort(arr []int, left,right int)  {
	mid := (left + right) >> 1
	subMergeSort(arr, left, mid)
	subMergeSort(arr, mid, right)

	merge()
}*/
/**
如何打印一个完全二叉树
---------------------------------------------------------------
0 1 2 3 4 5 6 7 8
 */
func printTree(arr []int, n,i,times int)  {
	// 用栈？？？？


	if i >= n - 1 {
		return
	}

	times++

	if i == 0 {
		// 打印树根
		fmt.Printf("    %d    ", arr[i])
		fmt.Println()
	}

	c1 := i * 2 + 1
	c2 := i * 2 + 2

	if c1 < n {
		// 打印左子树
		fmt.Printf("    %d    ", arr[c1])
	}
	if c2 < n {
		// 打印右子树
		fmt.Printf("    %d    ", arr[c2])

		fmt.Printf("(%d)(%d)(%d)", c2, times, power(2, times + 1) - 2)
		// 换行
		if power(2, times + 1) - 2 == c2 {
			println()
		}
	}

	printTree(arr, n, c1, times)
	printTree(arr, n, c2, times)
}


func power(n,m int) int {
	for i := 1;i < m;i++ {
		n *= n
	}

	return n
}

func heapSort(arr []int, len int) []int {
	maxParent := (len - 1 - 1) / 2
	for i := maxParent; i >= 0 ; i-- {
		heapify(arr, len,i)
	}
	//heapify(arr, len,0)

	for i := len - 1 ; i >= 0 ; i -- {
		arr[0],arr[i] = arr[i],arr[0]

		heapify(arr, i, 0)
	}

	return arr
}

func heapify(arr []int, n,i int) {
	if i >= n {
		return
	}

	c1 := 2 * i + 1
	c2 := 2 * i + 2

	maxIndex := i

	if c1 < n && arr[c1] > arr[maxIndex] {
		maxIndex = c1
	}

	if c2 < n && arr[c2] > arr[maxIndex] {
		maxIndex = c2
	}

	if maxIndex != i {
		swap(arr, maxIndex, i)
		heapify(arr, n, maxIndex)
	}
}

func swap(arr []int, i,j int)  {
	arr[i],arr[j] = arr[j],arr[i]
}

/*
----------------------------------------------------------------------------
 */

func shellSort(arr []int, len int) []int {
	// 先算出步长数组
	step := len
	if (step >> 1) > 0 {
		step = step >> 1
	}

	for step > 0 {
		for col := 0; col < step; col ++ {
			for begin := col + step; begin < len ; begin += step {
				cur := begin
				for cur > col && arr[cur] < arr[cur-step] {
					arr[cur], arr[cur-step]= arr[cur-step], arr[cur]
					cur -= step
				}
			}
		}
		//fmt.Println(step)
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