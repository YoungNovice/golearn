package main

import (
	"fmt"
	"retriever/mock"
)

type Retriever interface {
	Get(url string) string
}

// 在实现接口的时候 如果方法指针接收者那么赋值需要取地址|
// 如果方法全是值接收者 都可以
func main() {
	var r Retriever
	r = mock.RetrieverImp{Content: "this is a mock Retriever"}
	fmt.Println(r.Get("sss"))
}
