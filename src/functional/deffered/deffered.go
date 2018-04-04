package main

import (
	"bufio"
	"errors"
	"fmt"
	"functional/fib"
	"os"
)

// deffer 处理资源
func WriteFile(fileName string) {
	//file, err := os.Create(fileName)
	file, err := os.OpenFile(fileName, os.O_EXCL|os.O_CREATE, 0666)
	err = errors.New("this is my errors")
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			fmt.Println(pathError.Op, pathError.Err, pathError.Path)
		} else {
			panic(err)
		}
		return
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
