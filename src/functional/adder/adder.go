package main

import "fmt"

// 函数类型别名
type iAdder func(int) (int, iAdder)

// 1.递归调用 2.函数式编程 3.闭包
func adder(x int) iAdder {
	return func(temp int) (int, iAdder) {
		total := x + temp
		return total, adder(total)
	}
}

func main() {
	xxx := adder(0)
	for i := 0; i < 10; i++ {
		var sum int
		sum, xxx = xxx(i)
		fmt.Println(sum)
	}
}
