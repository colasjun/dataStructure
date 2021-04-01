package main

import (
	"fmt"
	"math/rand"
)

//插入排序
func main() {
	// 先随机生成一个int 数组
	len := 10
	arr := getArr(len)

	fmt.Printf("排序前%v\n",arr)

	var tmp int
	// 最外层是拿每一个数组的元素 内层是和已经排好序的数组依次比较
	for i := 1;i < len;i++ {
		//fmt.Printf("i=%d排序后%v\n",i,arr)
		for j := i; j >= 0 ;j-- {
			//fmt.Printf("j=%d排序后%v\n",j,arr)
			if arr[j+1] < arr[j] {
				tmp = arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
			}
		}

	}

	fmt.Printf("排序后%v",arr)
}

func getArr(num int) []int {
	var arr []int
	var i = 0
	for  i < num {
		arr = append(arr, rand.Intn(100))
		i ++
	}
	return arr
}
