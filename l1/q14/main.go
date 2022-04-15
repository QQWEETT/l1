package main

import "fmt"

func typeOf(v interface{}) {
	switch c := v.(type) {
	case int:
		fmt.Println("Int:", c)
	case string:
		fmt.Println("String:", c)
	case bool:
		fmt.Println("Bool:", c)
	case chan int:
		fmt.Println("Chan:", c)
	}
}

func main() {
	str := "Hello"
	num := 42
	boolean := true
	ch := make(chan int)
	typeOf(str)
	typeOf(num)
	typeOf(boolean)
	typeOf(ch)
}
