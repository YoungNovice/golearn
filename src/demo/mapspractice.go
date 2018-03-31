package demo

import "fmt"

func getStringMaxSubString(s string) (int)  {
	// 用map 记录每一个字符最后出现的位置
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if value, ok := lastOccurred[ch]; ok && value >= start {
			start = lastOccurred[ch] + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength

}

func main() {

	//utf8包
	// 关于字符串可以在strings包下

	fmt.Println(getStringMaxSubString("bbdsd"))

}
