package main

import (
	"bufio"
	"fmt"
	"functional/fib"
	"os"
)

// deffer 处理资源
func WriteFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	writer := bufio.NewWriter(file)

	defer writer.Flush()
	fib := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, fib())
	}
}

func main() {
	WriteFile("test.txt")
}
