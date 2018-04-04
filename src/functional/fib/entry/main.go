package main

import (
	"fmt"
	"functional/fib"
)

func main() {
	f := fib.Fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	f1 := fib.Fibonacci()
	fmt.Println(f1())
}
