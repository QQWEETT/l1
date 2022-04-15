package main

import (
	"fmt"
	"sync"
	"time"
)

type Data struct {
	sync.Mutex
	item map[int]int
}

func NewData() *Data {
	return &Data{
		item: make(map[int]int),
	}
}

func call(d *Data, subscriberId int) {
	for i := 0; i < 3; i++ {
		d.Lock()
		d.item[i] = subscriberId
		d.Unlock()
		fmt.Printf("The subscriber %v called on line %v \n", subscriberId, i)

	}
}

func main() {
	data := NewData()
	for i := 0; i < 3; i++ {
		go call(data, i)
	}
	time.Sleep(2 * time.Second)

	fmt.Println()
	data.Lock()
	for key, value := range data.item {
		fmt.Printf("ключ: %v значение: %v\n", key, value)
	}
	data.Unlock()
}
