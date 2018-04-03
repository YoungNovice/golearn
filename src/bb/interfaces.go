package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}
type Student struct {
	Human
	scholl string
	loan   float32
}
type Employee struct {
	Human
	company string
	money   float32
}
// Human 实现了SayHi 方法
func (h Human) SayHi()  {
	fmt.Printf("Hi , i am %s you can call me on %s\n", h.name, h.phone)
}
func (h Human) String() string  {
	return "Human"
}
// Human 实现了Sing 方法
func (h Human) Sing(lyrics string)  {
	fmt.Println("La la la la demaziya...", lyrics)
}
// 重写Human 的SayHi 方法
func (e Employee) SayHi ()  {
	fmt.Println("Hi, I am %s, I work at %s, Call me on %s\n",
		e.Human.name, e.company, e.phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike", 25, "2222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"paul", 26, "1111-222-XXX"}, "Hrvard", 1.00}
	sam := Employee{Human{"Sam", 36, "1111-222-XXX"}, "Golang Inc", 1000}
	tom := Employee{Human{"tom", 36, "1111-222-XXX"}, "Things Ltd", 5000}

	var men Men
	men = mike
	fmt.Println("This is mike, a Student:")
	men.SayHi()
	men.Sing("a diao")

	x := make([]Men, 3)
	x[0] = paul
	x[1]= sam
	x[2] = tom

}



