package main

import "fmt"

//牛客101
func main() {
	list := []int{1, 4, 3, 5, 2, 9, 0, 10, 7, 100}

	// list1 := []int{3, 4, 5}
	list = HeapSort(list)
	fmt.Print(list)
}
