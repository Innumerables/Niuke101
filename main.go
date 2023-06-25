package main

import "fmt"

//牛客101
func main() {
	list := []int{1, 4, 3, 5, 2, 9, 0, 10, 7, 100}
	list = QuickSort(list)
	fmt.Print(list)
}
