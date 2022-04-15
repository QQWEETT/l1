package main

import (
	"fmt"
	"time"
)

func say(ch chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		ch <- i
	}
	close(ch)
}

//func selectOne(data, exit chan int) {
//	x := 0
//	for {
//		select {
//		case data <- x:
//			x += 1
//		case <-exit:
//			fmt.Println("exit")
//			return
//		}
//		time.Sleep(time.Millisecond * 100)

//	}
//}

func main() {
	ch := make(chan int)
	go say(ch)
	for i := range ch {
		fmt.Println(i)
	}

	//	data := make(chan int)
	//	exit := make(chan int)
	//	go func() {
	//		for i := 0; i < 10; i++ {
	//			fmt.Println(<-data)
	//		}
	//		exit <- 0
	//	}()
	//	selectOne(data, exit)

	//	quit := make(chan bool)
	//	go func() {
	//		for {
	//			select {
	//			case <-quit:
	//				return
	//			default:
	//			}
	//		}
	//	}()
	//	quit <- true
}
