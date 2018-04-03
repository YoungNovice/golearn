package main

import (
	"io/ioutil"
	"fmt"
)

// ioutil 包可以方便的读取文件
func main() {
	readFile()
}



// if
func ifOne(x int) int {
	if x > 100 {
		return 100
	}else if x < 0 {
		return 0
	} else {
		return x
	}
}


func readFile()  {
	a := "abc"
	if bytes, e := ioutil.ReadFile(a); e != nil {
		fmt.Print(e)
	} else {
		fmt.Printf("%s \n", bytes)
	}
}