package main

import (
	"fmt"
	"retriever/mock"
	"retriever/real"
)

type Retriever interface {
	Get(url string) string
}

// 在实现接口的时候 如果方法指针接收者那么赋值需要取地址|
// 如果方法全是值接收者 都可以
func main() {
	var r Retriever
	r = mock.RetrieverImp{Content: "this is a mock Retriever"}
	inspect(r)
	r = &real.Retriever{}
	inspect(r)

	// type assertion
	if retriever, ok := r.(*real.Retriever); ok {
		fmt.Println(retriever.User)
	} else {
		fmt.Println("not real.Retriever	")
	}
	//fmt.Println(r.Get("http://www.baidu.com"))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.RetrieverImp:
		fmt.Println("contents:", v.Content)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.User)

	}
}
