package main

import (
	"fmt"
	"math"
)

func convertToByte(number int64) []int64 {
	bit := []int64{}
	for number != 1 {
		a := number % 2
		bit = append(bit, a)
		number /= 2
	}

	bit = append(bit, 1)
	for i, j := 0, len(bit)-1; i < j; i, j = i+1, j-1 {
		bit[i], bit[j] = bit[j], bit[i]
	}

	return bit
}

func add(number []int64, position int) []int64 {
	number[position] += 1
	return number
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func convertToNumber(number []int64) int64 {

	up := 0

	var a int64 = 0

	for i := len(number) - 1; i > -1; i-- {
		a += number[i] * int64(powInt(2, up))
		up++

	}
	return a
}

func main() {
	a := convertToByte(64)
	fmt.Println(a)
	s := add(a, 3)
	fmt.Println(s)
	d := convertToNumber(s)
	fmt.Println(d)
}
