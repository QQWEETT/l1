package main

import "fmt"

func getGroup(number []float64) map[int][]float64 {
	groups := make(map[int][]float64)
	for _, val := range number {
		group := int(val/10) * 10
		groups[group] = append(groups[group], val)
	}
	return groups

}

func main() {
	slice := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -15, -17, -9, 1}
	a := getGroup(slice)
	fmt.Println(a)
}
