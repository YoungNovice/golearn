package demo

import (
	"fmt"
	"os"
)

/**
	整数转换成 二进制数字
 */
/*
func convertToBin(n int) string  {

}
*/

func main() {
	sum := 0
	for i :=1; i<=100 ; i++ {
		sum += i
	}
	fmt.Println(sum, 1)
	file, e := os.Open("array sd")
	fmt.Println(file, e)

	// 就是 这里的for 就是while
	/*for contition {

	}*/


	// 什么条件都没有就是死循环
	/*for {

	}*/

}
