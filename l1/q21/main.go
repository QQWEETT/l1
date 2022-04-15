package main

import "fmt"

type cat interface {
	meow()
	run()
}

type meower struct{}

func (c *meower) meow() {
	fmt.Println("Meow")
}

func (c *meower) run() {
	fmt.Println("I run")
}

type dog interface {
	woof()
	run()
}

type doggy struct{}

func (d *doggy) woof() {
	fmt.Println("Woof")
}
func (d *doggy) run() {
	fmt.Println("I am running")
}

// Получаем ссылку на объект, который адаптируем
type dogAdapter struct {
	dog dog
}

func newDogAdapter(d dog) *dogAdapter {
	return &dogAdapter{
		dog: d,
	}
}

//Реализуем интерфейс, к которому адаптируемся
func (d *dogAdapter) meow() {
	d.dog.woof()
}
func (d *dogAdapter) run() {
	for i := 0; i < 5; i++ {
		d.dog.run()

	}
}

func main() {
	// создаем кота и собаку
	cat := &meower{}
	dog := &doggy{}
	//Оборачиваем собаку в адаптер
	dogAdapter := newDogAdapter(dog)

	fmt.Println("Cat says")
	testCat(cat)
	//Передача собаки за кошку тому же методу
	fmt.Println("The dog adapter says")
	testCat(dogAdapter)
}

func testCat(c cat) {
	c.meow()
	c.run()
}
