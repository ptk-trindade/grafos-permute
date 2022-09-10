package main

import (
	"fmt"
	"time"
)

// Print the [list] in the expected format
func show(list []int) {
	fmt.Print("{ ")
	lastItem := len(list) - 1
	for i := 0; i < lastItem; i++ {
		fmt.Print(list[i], ", ")
	}
	fmt.Print(list[lastItem], " ")
	fmt.Println("}")
}

// Permute the elements of the [list] between [index] and [length]
func permute(list []int, index int, length int) {
	if index == length {
		show(list)
	}
	for i := index; i < length; i++ {
		list[index], list[i] = list[i], list[index]
		permute(list, index+1, length)
		list[index], list[i] = list[i], list[index]
	}
}

func main() {
	var n int
	fmt.Scan(&n) // Get user input

	list := make([]int, n) // allocate a slice of size n
	for i := 0; i < n; i++ {
		list[i] = i + 1
	}

	start_time := time.Now()
	permute(list, 0, n)
	end_time := time.Now()

	fmt.Println("Time taken: ", end_time.Sub(start_time))
}
