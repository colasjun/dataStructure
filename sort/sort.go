package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	arr := getArr(25)

	fmt.Println("冒泡排序前", arr)
	sort := bubbleSort(arr, len(arr))
	fmt.Println("冒泡排序后", sort)
}

func bubbleSort (arr []int, len int) []int {
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