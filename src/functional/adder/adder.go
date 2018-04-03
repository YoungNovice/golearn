package adder

import "fmt"

/**
 func 函数有局部变量v 和 一个自由变量v
自由变量不是在函数体里面定义的， 他是函数所处的环境


*/

func adder() func(i int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(i int) (int, iAdder)

func adder2(x int) iAdder {
	return func(temp int) (int, iAdder) {
		return x + temp, adder2(x + temp)
	}
}

func main() {
	var add = adder()
	for i := 0; i < 10; i++ {
		fmt.Println(add(i))
	}

	adder := adder2(0)
	for i := 0; i < 10; i++ {
		var sum int
		sum, adder = adder(i)
		fmt.Println(sum)
	}
}
