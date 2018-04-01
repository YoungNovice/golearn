package main

import (
	"queue"
	"fmt"
)

func main() {
	q := queue.Queue{1}
	q.Push(1)
	fmt.Println(q.IsEmpty())
	pop := q.Pop()
	fmt.Println(pop)
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
