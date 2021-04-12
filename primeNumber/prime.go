package main

import "fmt"

func main()  {
	n := 10000
	// 暴力破解
	num := getPrimeNum(n)
	fmt.Printf("暴力法:前%d个数中素数的个数为%d个 \n", n, num)

	// 埃氏筛选法
	num = eratosthenes(n)
	fmt.Printf("埃筛法:前%d个数中素数的个数为%d个 \n", n, num)
}

//暴力法
func getPrimeNum(n int) int {
	var count = 0
	if n <= 1 {
		return count
	}

	for i := 2 ; i <= n ; i ++ {
		if isPrime(i) {
			count++
		}
	}

	return count
}

// 挨筛法
func eratosthenes(n int) int {
	var count = 0
	if n <= 1 {
		return count
	}

	// 定义map标记是否素数
	m := make(map[int]bool, n)  // false代表素数
	for i := 2 ; i <= n ; i ++ {
		if !m[i] {
			count++
			for j := i * i; j <= n; j += i {
				m[j] = true
			}
		}
	}

	return count
}

func isPrime(x int) bool {
	for i := 2;i * i  <= x;i++ {
		if x % i == 0 {
			return false
		}
	}

	return true
}