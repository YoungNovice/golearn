package demo

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// apply 函数接受一个f 函数作为参数  和两个int 函数 返回值是int
func apply(f func(a, b int) int, a, b int) int {
	// 返回f 函数的指针
	of := reflect.ValueOf(f).Pointer()
	// 在运行是获取f 函数的指针
	pc := runtime.FuncForPC(of)
	// 获取函数的名字
	funcName := pc.Name()
	fmt.Printf("calling function %s with args "+
		"(%d, %d)", funcName, a, b)
	return f(a, b)
}

// math 的pow函数没有重载形式， 我们可以给它写一个重载的int 处理
func pow(a, b int) int {
	pow := math.Pow(float64(a), float64(b))
	return int(pow)
}

func main() {

	apply(pow, 2, 3)

}
