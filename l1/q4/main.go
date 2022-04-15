package main

import (
	"fmt"
	"time"
)

const workers = 5

func workerMine(workerNum int, in <-chan interface{}) {

	for input := range in {
		fmt.Printf("Worker: %v | %v \n", workerNum, input)

	}
	fmt.Printf("Worker %v finished work \n", workerNum)
}

func main() {

	worker := make(chan interface{})
	for i := 0; i < workers; i++ {
		go workerMine(i, worker)
	}
	data := []interface{}{
		"Mithril", "Diamond", "Steel", "Iron", "Copper", "Coal", "Rock", "Brilliant", "Gold", "Tin",
		"aluminum", "Titanium",
	}
	for _, v := range data {
		worker <- v
	}
	close(worker)
	time.Sleep(time.Millisecond * 10)
}
