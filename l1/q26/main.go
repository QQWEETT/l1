package main

import (
	"fmt"
	"strings"
)

func unique(a string) bool {
	a = strings.ToLower(a) // преобразуем строку в нижний регистр
	for i := 0; i < len(a); i++ {
		res := strings.Count(a, string(a[i])) // считаем количество букв в строке
		if res > 1 {
			return false
			break
		} else {
			continue
		}
	}
	return true

}
func main() {
	letters := []string{"abcd", "abCdefAaf", "aabcd"}
	for i := range letters {
		fmt.Printf("%s - %v\n", letters[i], unique(letters[i]))
	}
}
