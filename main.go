package main

import "fmt"

func main() {
	fmt.Println("starting ...")
	/* Haystack */
	arr := []int{1, 6, 9, 55, 66, 74, 89, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116}
	fmt.Println(binarySearch(112, arr))

}

func binarySearch(needle int, haystack []int) bool {
	low := 0
	high := len(haystack) - 1
	step := 0
	fmt.Printf("total len is : %d\n------\n", high)

	for low <= high {
		step++
		median := (low + high) / 2
		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
		fmt.Printf("step %d:\n", step)
		fmt.Printf("low is : %d\nhigh is %d\n---------\n", low, high)
	}
	if low == len(haystack) || haystack[low] != needle {
		return false
	}
	return true
}
