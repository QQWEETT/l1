package main

import "fmt"

func set(s []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range s {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	setArr := set(arr)

	fmt.Printf("Arr: %v\nSet: %v\n", arr, setArr)
}
