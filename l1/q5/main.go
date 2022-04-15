package main

import (
	"fmt"
	"time"
)

func Read(ch chan string) {
	sec := 1
	for data := range ch {
		fmt.Println(data, "|", sec, " second(s)")
		sec += 1
	}
	fmt.Println("closed")
}

func main() {
	ch := make(chan string)
	seconds := 2
	go Read(ch)
	duration := time.Second * time.Duration(seconds)
	//устанавливаем таймер работы
	timer := time.NewTimer(duration)

	for {
		select {
		case <-timer.C:
			// по истечению таймера закрывается канал
			close(ch)
			fmt.Println("Time's out")
			time.Sleep(time.Millisecond * 1000)
			return
		default:
			ch <- "Hello!"
		}
		time.Sleep(time.Millisecond * 1000)
	}
}
