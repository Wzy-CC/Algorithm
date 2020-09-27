package main

import (
	"log"
)

// >>>>>>>>>>>>>>>>>> 排序算法

// 冒泡排序
// func bubbleSort(a []int) []int {
// 	// 每次排序前先进行深拷贝
// 	var results = make([]int, len(a))
// 	copy(results, a)
// 	// bubbleSort
// 	for i := 0; i < len(results)-1; i++ {
// 		for j := 0; j; j++ {
// 			if awf {
// 				// reverse
// 			}
// 		}
// 	}
// 	return results
// }

// 插入排序

// 选择排序

// 希尔排序

// 快速排序
func quickSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	var sentinel = a[0]
	left := 1           // left index
	right := len(a) - 1 // right index
	for left < right {
		for a[right] >= sentinel && left < right {
			right--
		}
		for a[left] <= sentinel && left < right {
			left++
		}
		if left < right {
			a[left], a[right] = a[right], a[left]
		}
		if left == right {
			break
		}
	}
	mid := left
	a[0], a[mid] = a[mid], a[0] // reverse

	quickSort(a[:mid])   // quick sort left
	quickSort(a[mid+1:]) // quick sort right
	return a
}

// 归并排序
func merge(a, b []int) (result []int) { // merge two array
	// a:l b:r
	l, r := 0, 0 // l,r is index of a b
	for l < len(a) && r < len(b) {
		if a[l] < b[r] {
			result = append(result, a[l])
			l++
		} else {
			result = append(result, b[r])
			r++
		}
	}

	result = append(result, a[l:]...) // else Ops
	result = append(result, b[r:]...) // else Ops
	// log.Println(">>>", result)        // test code
	return result
}

func mergeSort(a []int) []int {
	alen := len(a)
	if alen <= 1 { // sp condition
		return a
	}

	mid := alen / 2
	left := mergeSort(a[:mid])
	right := mergeSort(a[mid:])
	return merge(left, right)
}

// 堆排序

// 线性排序算法

// 自省排序

// 计数排序

// 桶排序

// >>>>>>>>>>>>>>>>>> 递归与分治

// >>>>>>>>>>>>>>>>>> 动态规划

// >>>>>>>>>>>>>>>>>> 贪心

// >>>>>>>>>>>>>>>>>> 回溯

// >>>>>>>>>>>>>>>>>> 搜索

// DFS

// BFS

func main() {
	array := []int{3, 2, 1, 3, 7, 634, 543, 24, 656, 25, 7, 17, 4632, 5346, 536, 6234, 45, 56, 4, 32, 14, 23, 547, 568, 5, 85, 4, 32, 214, 46, 54, 56}
	// log.Println(bubbleSort(array)) // 冒泡排序算法
	log.Println(quickSort(array)) // quick sort
	// log.Println(mergeSort(array)) // merge sort
}
