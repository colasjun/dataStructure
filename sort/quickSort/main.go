package main

import (
	"fmt"
	"math/rand"
)

// 快速排序算法
func main() {
	// 先随机生成一个int 数组
	len := 10
	arr := getArr(len)

	fmt.Printf("排序前%v",arr)



	quickSort(arr, 0, len - 1)

	fmt.Printf("选择排序后%v",arr)
}


func quickSort(theArray []int, start int, end int)[]int {
	if (start<end){
		m, n := start, end
		base := theArray[m]
		for {
			if (m < n){
				for{
					if((m < n)&&(theArray[n]>=base)){ // 大于基数的数字不动 右边指针向左移动
						n--
					}else{
							// 小于基数时 基数位置存放当前值
						theArray[m] = theArray[n]
						break
					}
				}
				for{
					//
					if((m < n)&&(theArray[m]<=base)){
						m++
					}else{
						theArray[n] = theArray[m]
						break
					}
				}
			}else{
				break
			}
			theArray[m] = base
			quickSort(theArray, start, m-1)
			quickSort(theArray, n+1, end)
		}
	}
	return theArray
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

