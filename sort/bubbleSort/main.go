package main

import (
	"fmt"
	"math/rand"
)

//冒泡排序
func main() {
	// 先随机生成一个int 数组
	len := 10
	arr := getArr(len)

	fmt.Printf("排序前%v",arr)

	/*var temp int
	for i := len - 1; i >= 0 ; i-- {
		for m := len - 1; m >= 1; m-- {
			if arr[m] <= arr[m-1] {
				temp = arr[m]
				arr[m] = arr[m-1]
				arr[m-1] = temp
			}
		}
	}*/

	/*n := 0
	for i := 0; i <= len - 1; i++ {
		for j := 0; j < len - 1 - n; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		n++
	}*/

	for i := 1; i <= len ;i++ {
		for j := 1; j <= len - i;j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}

	fmt.Printf("排序后%v",arr)
}

func getArr(num int) []int {
	var arr []int
	var i = 0
	for  i < num {
		arr = append(arr, rand.Intn(1000))
		i ++
	}
	return arr
}
