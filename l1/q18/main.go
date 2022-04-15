package main

import (
	"fmt"
	"sync"
)

// Создаем counter для параллельного подсчета
type Counter struct {
	mu      sync.Mutex
	counter int
}

// Создаем конструктор счетчика
func NewCounter() *Counter {
	return &Counter{}
}

//Увеличиваем счетчик на 1
func (c *Counter) Incr() {
	c.mu.Lock()
	c.counter++
	defer c.mu.Unlock()

}

//Печатает string
func (c *Counter) String() string {
	return fmt.Sprint(c.counter)
}

func main() {
	counter := NewCounter()
	var wg sync.WaitGroup
	loop := 500

	for i := 1; i <= loop; i++ {
		wg.Add(1)

		go func(c int) {
			defer wg.Done()
			counter.Incr()
		}(i)
	}

	wg.Wait()
	fmt.Println("Final count = ", counter)
}
