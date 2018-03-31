package main

import "fmt"

// 数组
// 数组是值类型 [3]int 和 [5]int 不是同一个type
func main() {

	//var arr1 [5]int
	//arr2 := [3]int{1,2,3}
	// 如果不大点就是切片
	//arr3 := [...]int{1,2,4}


	a := 1
	change(&a)
	fmt.Println(a)
}
func change(a *int) {
	*a = 2
}
