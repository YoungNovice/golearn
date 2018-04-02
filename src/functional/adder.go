package main

import "fmt"

func adder() func(i int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	var add = adder()
	for i := 0; i < 10; i++ {
		fmt.Println(add(i))
	}
}
