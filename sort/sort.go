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
	for i := 1; i < len(input)-1; i++ {
		for j := i; j > 0; j-- {
			if input[j] < input[j-1] {
				input[j-1], input[j] = input[j], input[j-1]
			}
		}
	}
	return input
}

// 归并排序
func MergingSort(input []int) []int {
	return input
}

// 快速排序
func QuickSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}
	var (
		left, right []int
	)
	mid := input[0]
	for i := 1; i < len(input); i++ {
		if input[i] > mid {
			right = append(right, input[i])
		} else {
			left = append(left, input[i])
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left, mid), right...)
}

// 堆排序
