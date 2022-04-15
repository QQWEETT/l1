package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1
	// Создаем точку "опоры"
	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]
	// сравниваем каждое число с точкой "опоры", если true, то a[left] станет a[i], a a[i] станет a[left]
	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	// свапаем местами
	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])   // разбиваем на подмассивы левее left и  сортируем
	quicksort(a[left+1:]) // разбиваем на подмассивы правее left и  сортируем

	return a
}

func main() {
	a := []int{3, 5, 6, 1, 8, 6, 3, 7, 9}
	fmt.Println(quicksort(a))
	sort.Ints(a)
	fmt.Println(a)
}
