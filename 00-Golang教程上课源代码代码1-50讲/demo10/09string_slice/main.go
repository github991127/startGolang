package main

import "fmt"

func main() {
	// s1 := "big"
	// byteStr := []byte(s1)
	// byteStr[0] = 'p'
	// fmt.Println(string(byteStr))

	s2 := "你好golang"
	runeStr := []rune(s2)
	fmt.Println(runeStr) //[20320 22909 103 111 108 97 110 103]
	runeStr[0] = '大'
	fmt.Println(string(runeStr))

}
