package main

import (
	"fmt"
	"strings"
)

// 生成相应的二维数组
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

// 牛客101
func main() {
	// list := []int{1, 4, 3, 5, 2, 9, 0, 10, 7, 100}
	// list = HeapSort(list)
	// fmt.Print(list)
	// l := findPeakElement(list)
	// fmt.Println(l)
	// a := []int{1, 2, 3, 4, 5, 6, 7, 0}
	// b := InversePairs(a)
	// fmt.Println(b)
	// demo := "1.01.01.21.00"
	// string_slice := strings.Split(demo, ".")

	// fmt.Println("result:", string_slice)
	// fmt.Println("len:", len(string_slice))
	// fmt.Println("cap:", cap(string_slice))
	// tt, _ := strconv.Atoi(string_slice[4])
	// fmt.Printf("%T,%v", tt, tt)
	// root := &TreeNode{
	// 	Val: 1,
	// 	Left: &TreeNode{
	// 		Val: 2,
	// 		Left: &TreeNode{
	// 			Val: 4,
	// 		},
	// 		Right: &TreeNode{
	// 			Val: 5,
	// 		},
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 3,
	// 	},
	// }
	// res := preorderTraversal(root)
	// fmt.Println(res)
	s := "jklasd"
	t := strings.Split(s, ",")
	fmt.Printf("%T", t)
}
