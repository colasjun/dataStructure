package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

func main()  {

	file, _ := os.Create("./fbnq.trace")

	defer file.Close()

	trace.Start(file)

	go func() {
		time.Sleep(time.Second * 2)
	}()

	// 使用迭代方法实现斐波那契数列
	data := getFibData(2)

	fmt.Printf("data:%d" , data)

	m := make(map[int]int)
	m[1] = 1

	a := 2
	if _, ok := m[10]; ok {
		a = 22
	}

	for i := 1; i < a; i++ {

	}

	for a < 10 {
		a ++
	}

	trace.Stop()
}

func getFibData(n int) int {
	arr :=[]int{0,1}

	if n <= 2 {
		return arr[n-1]
	}

	firstNum := arr[0]
	SecondNum := arr[1]
	ThirdNum := 0

	for i := 3 ; i <=  n ; i++ {
		ThirdNum = firstNum + SecondNum
		firstNum = SecondNum
		SecondNum = ThirdNum
	}

	return ThirdNum
}


