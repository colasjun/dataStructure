package main

import (
	"fmt"
	"math/rand"
)

//选择排序
func main() {
	// 先随机生成一个int 数组
	len := 17
	arr := getArr(len)

	fmt.Printf("排序前%v",arr)

	/*var max int
	var temp int
	var maxKey int
	for i := 0;i < len;i++ {
		for j := len - i;j >= 0;j-- {
			if arr[j] > max {
				max = arr[j]
				maxKey = j
			}
		}

		if max >= arr[len-i] {
			temp = arr[maxKey]
			arr[maxKey] = arr[len-i]
			arr[len-i] = temp
		}

		max,temp,maxKey = 0,0,0
	}*/

	max := 0
	maxKey := 0
	for i := 0; i < len; i++ {
		for j := 0; j < len - i ; j++ {
			if arr[j] > max {
				max = arr[j]
				maxKey = j
			}

			//fmt.Println(j, max)

		}
		//return
		// 在放入位置
		arr[len-i-1],arr[maxKey] = arr[maxKey],arr[len-i-1]
		max = 0
		maxKey = 0
	}

	fmt.Printf("\n选择排序后%v\n",arr)
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
