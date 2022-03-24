package myMath

import "fmt"

func gcd(x, y int) int {
	for y != 0 {
		fmt.Printf("%d %d %d\n", x, y, x%y)
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func mySlice() {
	var m, n int
	// 切片 []不指定大小
	data := []int{1, 2, 3, 4, 5}
	mydata := data[m:n]
	r := make([]int, len(mydata))
	// 大的切片 data 缩小成小的切片 mydata
	copy(r, mydata)

}
