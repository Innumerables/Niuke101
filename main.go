package main

import "fmt"

//生成相应的二维数组
func createArry(l, m int) [][]int {
	arry := [][]int{}
	for i := 0; i < l; i++ {
		tmp := []int{}
		for j := 0; j < m; j++ {
			tmp = append(tmp, j)
		}
		arry = append(arry, tmp)
	}
	return arry
}

//牛客101
func main() {
	list := []int{1, 4, 3, 5, 2, 9, 0, 10, 7, 100}
	// list = HeapSort(list)
	// fmt.Print(list)
	l := findPeakElement(list)
	fmt.Println(l)
	a := []int{1, 2, 3, 4, 5, 6, 7, 0}
	b := InversePairs(a)
	fmt.Println(b)
}
