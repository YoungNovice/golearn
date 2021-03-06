package fib

import (
	"fmt"
	"io"
	"strings"
)

type intGen func() int

func (i intGen) Read(p []byte) (n int, err error) {
	// 获取数列的下一个数字
	next := i()
	if next > 10000 {
		return 0, io.EOF
	}
	// 获取这个字符串
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func Fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
