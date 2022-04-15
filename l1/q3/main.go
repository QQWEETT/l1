package main

import "fmt"

const sum = 0

func SumSquares(nums []int) int {
	ch := make(chan int)
	var sum int

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
		sum += val
	}
	return sum
}
func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	sumSq := SumSquares(numbers[:])
	fmt.Printf("Сумма квадратов чисел = %d ", sumSq)
}
