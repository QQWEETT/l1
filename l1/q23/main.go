package main

import "fmt"

//долгий вариант
func remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...) // сдвигаем все элементы справа от индекса удаления на один влево
}

//быстрый вариант
func remove1(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1] // переставляем i-ый элемент слайса в конец
	return slice[:len(slice)-1]    // возвращаем n-1 первых элементов
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(arr)
	fmt.Println(remove1(arr, 3))

	//fmt.Println(remove(arr, 3))
}
