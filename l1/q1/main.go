package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

type Action struct {
	Human
	School string
}

func (h *Action) Refresh(name string, age int, school string) {
	h.Name = name
	h.Age = age
	h.School = school
}

func newAction(name string, age int, school string) *Action {
	return &Action{
		Human{
			Name: name,
			Age:  age,
		},
		school,
	}
}
func main() {
	s1 := Action{Human{Name: "Andrew", Age: 20}, "School1"}
	fmt.Println(s1)
	s1.Refresh("Nikita", 25, "School2")
	fmt.Println(s1)

	s2 := newAction("George", 20, "School1")
	fmt.Println(s2)
	s2.Refresh("George", 15, "School2")
	fmt.Println(s2)
}
