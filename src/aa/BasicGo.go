package main

import (
	"fmt"
	"math/cmplx"
)

func varibale() {
	var a int
	var s string = "sss"
	fmt.Printf("%d %q\n", a, s)
}

func variableBsd() {
	var a, b string = "1", "2"
	c, d := 1, 2
	fmt.Println(a, b, c, d)
}

func enloer()  {

	 c := 3 + 4i
	 fmt.Println(cmplx.Abs(c))

}

func main() {
	fmt.Println("Hello World")
	varibale()
	enloer()
}
