package main

// 判断interface 类型里面存储的数据
import (
	"fmt"
	"strconv"
)
// 定义Element 它是一个interface{} 可以存放任意类型
type Element interface {}
// List 是一个[]interface{} 切片
type List []Element

type Person struct {
	name string
	age int
}

func (p Person) String() string  {
	return "(name:) " + p.name + "- age: " + strconv.Itoa(p.age) + "years) "
}

func main() {
	list := make([]Element, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"yangxuan", 25}

	// while
	for index, element := range list {
		if value, ok := element.(string); ok {
			fmt.Println(value, "index:= " , index)
		}
	}

	
}
