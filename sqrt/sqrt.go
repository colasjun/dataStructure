package main

import (
	"fmt"
	"math"
)

// 在不使用自带函数 求x的平方根的整数部分

// 考察 二分法 牛顿迭代

func main()  {

	// 二分
	fmt.Println(getSqrtBinarySearch(25))

	// 牛顿
	fmt.Println(newton(5))
}

func getSqrtBinarySearch(n int) int {
	index := -1
	left := 0
	right := n
	for left <= right {
		middle := left + (right - left) / 2
		if (middle * middle) <= n {
			index = middle
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return index
}

func newton(n int) string {
	if n == 0 {
		return  "0"
	}

	C, x0 := float32(n), float32(n)
	for {
		xi := 0.5 * (x0 + C/x0)
		if math.Abs(float64(x0 - xi)) < 1e-7 {
			break
		}
		x0 = xi
	}

	return fmt.Sprintf("%.4f", x0)
}