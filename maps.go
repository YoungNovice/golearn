package main

import "fmt"

func main() {
	m1 := map[string]string{"name":"yangxuan", "name2": "qulianghong"}
	var m2 map[string]int
	m3 := make(map[string]int, 3)
	m3["a"] = 3
	fmt.Println(m1, m2, m3)

	for k, v := range m3 {
		fmt.Println(k, v)
	}
}
