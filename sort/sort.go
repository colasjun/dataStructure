package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	arr := getArr(100000)
	fmt.Printf("排序时间比较.元素一共%d 其中前面3个元素为%d %d %d", len(arr), arr[0], arr[1], arr[2])

	arr1 := getSameArr(arr)
	arr2 := getSameArr(arr)
	arr3 := getSameArr(arr)
	arr4 := getSameArr(arr)
	arr5 := getSameArr(arr)
	arr6 := getSameArr(arr)
	//arr7 := arr

	hr()
	//fmt.Println("冒泡排序前", arr)
	start := time.Now().UnixNano()  / 1e6
	fmt.Printf("冒泡其中前面3个元素为%d %d %d \n", arr1[0], arr1[1], arr1[2])
	sort := bubbleSort(arr1, len(arr1))
	fmt.Printf("冒泡排序执行时间:%d 毫秒 \n",  time.Now().UnixNano() / 1e6 - start)
	verifySort(arr1, sort)
	//fmt.Println("冒泡排序后", sort)
	hr()
	//fmt.Println("选择排序前", arr)
	start = time.Now().UnixNano()  / 1e6
	fmt.Printf("选择其中前面3个元素为%d %d %d \n", arr2[0], arr2[1], arr2[2])
	sort = selectSort(arr2, len(arr2))
	fmt.Printf("选择排序执行时间:%d 毫秒 \n",  time.Now().UnixNano() / 1e6 - start)
	verifySort(arr2, sort)
	//fmt.Println("选择排序后", sort)
	hr()
	start = time.Now().UnixNano()  / 1e6
	fmt.Printf("插入其中前面3个元素为%d %d %d \n", arr3[0], arr3[1], arr3[2])
	//fmt.Println("插入排序前", arr3)
	sort = insertSort(arr3, len(arr3))
	fmt.Printf("插入排序执行时间:%d 毫秒 \n", time.Now().UnixNano() / 1e6 - start)
	verifySort(arr3, sort)
	//fmt.Println("插入排序后", sort)
	hr()
	start = time.Now().UnixNano()  / 1e6
	fmt.Printf("希尔其中前面3个元素为%d %d %d \n", arr4[0], arr4[1], arr4[2])
	//fmt.Println("希尔排序前", arr4)
	sort = shellSort(arr4, len(arr4))
	fmt.Printf("希尔排序执行时间:%d 毫秒 \n",  time.Now().UnixNano() / 1e6 - start)
	verifySort(arr4, sort)
	//fmt.Println("希尔排序后", sort)
	hr()
	//arr = getArr(25)
	start = time.Now().UnixNano()  / 1e6
	fmt.Printf("堆其中前面3个元素为%d %d %d \n", arr5[0], arr5[1], arr5[2])
	//fmt.Println("堆排序前", arr)
	sort = heapSort(arr5, len(arr5))
	fmt.Printf("堆排序执行时间:%d 毫秒 \n",  time.Now().UnixNano() / 1e6 - start)
	verifySort(arr5, sort)
	//fmt.Println("堆排序后", sort)
	hr()
	start = time.Now().UnixNano()  / 1e6
	fmt.Printf("快排其中前面3个元素为%d %d %d \n", arr6[0], arr6[1], arr6[2])
	//arr = getArr(25)
	//fmt.Println("快排序前", arr)
	sort = quickSort(arr6, 0, len(arr6) - 1)
	fmt.Printf("快排执行时间:%d 毫秒 \n",  time.Now().UnixNano() / 1e6 - start)
	verifySort(arr6, sort)
	//fmt.Println("快排序后", sort)
}

func getSameArr(arr []int) []int  {
	re := make([]int, len(arr))
	copy(re, arr)
	return re
}

/**
----------------------------------------------------------------
 */
func quickSort(arr []int, start int, end int) []int {
	// 递归条件
	if start < end {
		// 轴点元素
		left := start
		right := end

		pointVal := arr[left]

		// 双指针
		for left < right {
			for left < right {
				if arr[right] < pointVal {
					arr[left] = arr[right]
					left++
					break
				} else {
					// 右边大于不动 指针--
					right--
				}
			}

			for left < right {
				if arr[left] >= pointVal {
					arr[right] = arr[left]
					right--
					break
				} else {
					// 右边大于不动 指针--
					left++
				}
			}
		}

		arr[left] = pointVal
		//fmt.Println(start, left,end)
		//return arr
		quickSort(arr, start, left-1)
		quickSort(arr, left+1, end)

	}
	return arr
}

/**
----------------------------------------------------------------
*/

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