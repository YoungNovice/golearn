package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func main() {
	print100(100)
	fmt.Println(printBin(1024))
	printfile("abc")
}

func printfile(filename string)  {
	file, err := os.Open(filename)
	if err == nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan()  {
		println(scanner.Text())
	}

}

func printBin(x int) string {
	result := ""
	for ; x > 0; x /= 2 {
		yushu := x % 2
		result = strconv.Itoa(yushu) + result
	}
	return result
}

func print100(x int) {
	sum := 0
	for i := 1; i < x; i++ {
		sum += i
	}
	fmt.Println(sum)
}
