package main

import (
	"fmt"
	"sort"
)

func binarySearch(nums []int, value int) int {
	start_index := 0
	end_index := len(nums) - 1
	// запускаем цикл до того момента, пока start_index не будет равен или меньше end_index
	for start_index <= end_index {

		median := (start_index + end_index) / 2
		if nums[median] < value {
			start_index = median + 1
		} else {
			end_index = median - 1
		}

	}

	if start_index == len(nums) || nums[start_index] != value {
		return -1
	} else {
		return start_index
	}

}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(items, 63))
	fmt.Println(sort.SearchInts(items, 63))
}
