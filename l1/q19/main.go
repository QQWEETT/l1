package main

import (
	"fmt"
	"strings"
)

func reverse1(s string) string {
	var a string
	ss := strings.Split(s, "") //преобразуем строку в слайс
	for i := len(ss) - 1; i != -1; i-- {
		a += ss[i] + ""

	}
	return a
}

func reverse2(s string) []string {
	ss := strings.Split(s, "") //преобразуем строку в слайс
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
	return ss
}

func main() {
	r1 := reverse1("гигарыба")
	fmt.Println(r1)
	r2 := reverse2("гигарыба")
	fmt.Println(strings.Join(r2, ""))
}
