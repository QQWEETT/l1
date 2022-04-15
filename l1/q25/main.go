package main

import (
	"fmt"
	"time"
)

func Sleep(t time.Duration) {
	<-time.After(t)
}

func main() {
	fmt.Println(1)
	Sleep(3 * time.Second)
	fmt.Println(3)
}
