package main

import (
	"fmt"
	"math/big"
)

// складывание
func add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

// вычитание
func sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

// умножение
func multiply(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

// деление
func divide(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func main() {

	a := big.NewInt(int64(1 << 30))
	b := big.NewInt(int64(1 << 25))

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("a + b = ", add(a, b))

	fmt.Println("a - b = ", sub(a, b))

	fmt.Println("a * b = ", multiply(a, b))

	fmt.Println("a / b = ", divide(a, b))

}
