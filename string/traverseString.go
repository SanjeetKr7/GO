package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcdefgh"
	k := []rune(s)
	var str []rune
	for i := 0; i < len(k); i++ {
		fmt.Println(string(k[i]))
		str = append(str, k[i])
	}
	fmt.Println(string(str))

	s1 := "abcdefghi"
	k1 := []rune(s1)

	fmt.Println(strings.Compare(string(k1[:len(k1)-1]), string(k)))
}
