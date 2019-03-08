package main

import "fmt"

func wordCheck(s string, i int) map[byte]int {
	count := map[byte]int{}
	for ; i < len(s); i++ {
		count[s[i]]++
	}
	return count
}

func wordCount(s string, i int) {
	count := 0
	for ; i < len(s); i++ {
		count++
	}
	fmt.Println(count)
}

func inlinePrint() {
	fmt.Println("123")
}

func main() {
	s := "abcd1234abcd1234"
	fmt.Println(wordCheck(s, 4))
	wordCount(s, 4)
	inlinePrint()
}
