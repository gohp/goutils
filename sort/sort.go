package sort

/**
* @Author: Jam Wong
* @Date: 2020/6/21
 */

// 冒泡排序
func BubbleSort(input []int) []int {
	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input)-1-i; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	return input
}

// 选择排序
func SelectSort(input []int) []int {
	for i := 0; i < len(input)-1; i++ {
		minIndex := i

		for j := i + 1; j < len(input); j++ {
			if input[j] < input[minIndex] {
				minIndex = j
			}
		}

		input[minIndex], input[i] = input[i], input[minIndex]
	}
	return input
}

// 插入排序
func InsertSort(input []int) []int {
	return input
}

// 归并排序
func MergingSort(input []int) []int {
	return input
}

// 快速排序
func QuickSort(input []int) []int {
	return input
}

// 堆排序
