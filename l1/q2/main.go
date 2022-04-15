package main

import (
	"fmt"
)

func Squares(nums []int) []int {
	ch := make(chan int) //создаем канал
	squares := make([]int, 0, len(nums))

	//Создаем горутину для подсчета квадрата каждого числа
	go func(ch chan int) {
		for _, v := range nums {
			ch <- v * v
		}
		close(ch)
	}(ch)

	for {
		val, err := <-ch
		if !err {
			break
		}
		fmt.Println(val)
		squares = append(squares, val)
	}
	return squares
}

func main() {

	numbers := [5]int{2, 4, 6, 8, 10}

	squares := Squares(numbers[:])

	for i := range numbers {
		fmt.Printf("Квадрат числа %d = %d\n", numbers[i], squares[i])
	}
}
