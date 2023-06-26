package main

import (
	"fmt"
	"math"
)

//十大排序算法
//名称      平均时间复杂度   最好情况     最坏情况    空间复杂度   稳定性

// 冒泡排序      O(n^2)        O(n)       O(n^2)        O(1)      稳定
// 主要思想：每次选出最大的一个放到后边，比较相邻的元素将大的元素后移
func BublbeSort(arr []int) []int {
	l := len(arr)
	for n := 0; n < l-1; n++ {
		flag := 0
		for i := 0; i < l-n-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				flag = 1
			}
		}
		if flag == 0 {
			break
		}
	}
	return arr
}

// 选择排序      O(n^2)        O(n^2)     O(n^2)        O(1)      不稳定
// 主要思想：从未排序序列中找出最大（小）元素放到序列起始位置，然后继续重复该步骤
func SelectSort(arr []int) []int {
	l := len(arr)
	for n := 0; n < l; n++ {
		min := arr[n]
		minIndex := n
		for i := n + 1; i < l; i++ {
			if arr[i] < min {
				min = arr[i]
				minIndex = i
			}
		}
		arr[n], arr[minIndex] = arr[minIndex], arr[n]
	}
	return arr
}

// 插入排序      O(n^2)        O(n)       O(n^2)        O(1)      稳定\
// 主要思想：将第一个看作待排序列，依次后移，将已过的序列排成有序序列，直到末尾
func InsertSort(arr []int) []int {
	l := len(arr)
	for n := 1; n < l; n++ {
		preIndex := n - 1
		mid := arr[n]
		for preIndex >= 0 && arr[preIndex] > mid {
			arr[preIndex+1] = arr[preIndex]
			preIndex -= 1
		}
		arr[preIndex+1] = mid
	}
	return arr
}

// 希尔排序      O(nlogn)   O(nlog^2 n)    O(nlog^2 n)  O(1)      不稳定
// 基本思想：先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，待整个序列中的记录“基本有序”时。再对全体记录依次进行直接插入排序
func ShellSort(arr []int) []int {
	l := len(arr)
	for gap := l / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < l; i++ {
			min := arr[i]
			j := i
			for ; j >= gap && min < arr[j-gap]; j = j - gap {
				arr[j] = arr[j-gap]
			}
			arr[j] = min
		}
	}
	return arr
}

// 归并排序      O(nlogn)      O(nlogn)    O(nlogn)     O(n)      稳定
// 采用分治法，过程是首先将序列拆分成子序列，对子序列进行排序，将排序好的子序列合并，得到一个有序的序列
func MergeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	mid := n / 2
	left := arr[:mid]
	right := arr[mid:]
	return merge(MergeSort(left), MergeSort(right))
}
func merge(left, right []int) []int {
	arr := []int{}
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			arr = append(arr, left[0])
			left = left[1:]
		} else {
			arr = append(arr, right[0])
			right = right[1:]
		}
	}
	if len(left) == 0 {
		arr = append(arr, right...)
	}
	if len(right) == 0 {
		arr = append(arr, left...)
	}
	return arr
}

// 快速排序      O(nlogn)      O(nlogn)     O(n^2)    O(logn)     不稳定
// 从数列中挑选一个元素作为基准，通过一趟排序将待排记录分为两部分，一部分比元素大，一部分比元素小，分别对这两部分继续进行排序，以使整个序列有序
func QuickSort(arr []int) []int {
	return Quick(arr, 0, len(arr)-1)
}
func Quick(arr []int, left int, right int) []int {
	if left < right {
		partitionIndex := Partition1Way(arr, left, right)
		Quick(arr, left, partitionIndex-1)
		Quick(arr, partitionIndex+1, right)
	}
	return arr
}

func Partition1Way(arr []int, left, right int) int {
	privot := left
	index := privot + 1
	for i := index; i <= right; i++ {
		if arr[privot] > arr[i] {
			arr[index], arr[i] = arr[i], arr[index]
			index += 1
		}
	}
	arr[privot], arr[index-1] = arr[index-1], arr[privot]
	return index - 1
}

// 堆排序        O(nlogn)      O(nlogn)     O(nlogn)      O(1)     不稳定
// 大顶堆（升序序列），小顶堆（降序序列）
// 主要思想：创建一个堆，把堆首（最大值）与堆尾交换，堆的尺寸减一，重复交换直到尺寸为1
func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
func HeapSort(arr []int) []int {
	arrlen := len(arr)        //以数组模拟堆
	buildMaxHeap(arr, arrlen) //排好序的大顶堆
	for i := arrlen - 1; i >= 0; i-- {
		swap(arr, 0, i) //将最大值交换到最后一位
		arrlen -= 1     //向前移
		heapify(arr, 0, arrlen)

	}
	return arr
}
func buildMaxHeap(arr []int, arrlen int) {
	for i := arrlen / 2; i >= 0; i-- {
		heapify(arr, i, arrlen) //
		fmt.Println(arr)
	}
}
func heapify(arr []int, i, arrlen int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < arrlen && arr[left] > arr[largest] {
		largest = left
	}
	if right < arrlen && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		swap(arr, i, largest)
		heapify(arr, largest, arrlen)
	}
}

// 计数排序      O(n+k)        O(n+k)       O(n^2)        O(k)      稳定
// 主要思想：计算存储空间，减去最小值将其存入指定位置，位置加最小值代表数据
func CountingSort(arr []int) []int {
	length := len(arr)
	maxValue, smallest := getMaxValue(arr) //计算所需存储的空间大小
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen)
	sortedIndex := 0
	// 统计每个元素出现的个数,取出的值减去最小值即为该元素所需要存在的位置
	for i := 0; i < length; i++ {
		bucket[arr[i]-smallest] += 1
	}
	// 按照统计结果写入arr
	for j := 0; j < length; j++ { //从小的位置依次取出即实现排序
		for bucket[j] > 0 {
			arr[sortedIndex] = j + smallest // j的位置就是已排序好的序列，将j加上最小值即是原始的值
			sortedIndex++
			bucket[j]--
		}
	}
	return arr
}

// 获得数组的最值差
func getMaxValue(arr []int) (int, int) {
	largest := math.MinInt32
	smallest := math.MaxInt32
	for i := 0; i < len(arr); i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
		if arr[i] < smallest {
			smallest = arr[i]
		}
	}
	maxValue := largest - smallest
	return maxValue, smallest
}

//桶排序        O(n+k)        O(n+k)     O(n^2)        O(n+k)      稳定
//基数排序      O(n*k)        O(n*k)      O(n*k)       O(n+k)      稳定
