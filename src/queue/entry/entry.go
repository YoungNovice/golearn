package main

import (
	"fmt"
	"queue"
)

func main() {
	q := queue.Queue{1}
	q.Push(1)
	q.Push("sd")
	fmt.Println(q.IsEmpty())
	pop := q.Pop()
	fmt.Println(pop)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
