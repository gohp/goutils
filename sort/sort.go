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
	if len(input) < 2 {
		return input
	}
	middle := len(input) / 2
	left := MergingSort(input[:middle])
	right := MergingSort(input[middle:])
	return merge(left, right)
}

func merge(a, b []int) []int {
	var result []int
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
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
func HeapSort(input []int) []int {

	N := len(input) - 1 //s[0]不用，实际元素数量和最后一个元素的角标都为N
	//构造堆
	//如果给两个已构造好的堆添加一个共同父节点，
	//将新添加的节点作一次下沉将构造一个新堆，
	//由于叶子节点都可看作一个构造好的堆，所以
	//可以从最后一个非叶子节点开始下沉，直至
	//根节点，最后一个非叶子节点是最后一个叶子
	//节点的父节点，角标为N/2
	for k := N / 2; k >= 1; k-- {
		sink(input, k, N)
	}
	//下沉排序
	for N > 1 {
		swap(input, 1, N) //将大的放在数组后面，升序排序
		N--
		sink(input, 1, N)
	}

	return input
}

//下沉（由上至下的堆有序化）
func sink(s []int, k, N int) {
	for {
		i := 2 * k
		if i > N { //保证该节点是非叶子节点
			break
		}
		if i < N && s[i+1] > s[i] { //选择较大的子节点
			i++
		}
		if s[k] >= s[i] { //没下沉到底就构造好堆了
			break
		}
		swap(s, k, i)
		k = i
	}
}

func swap(s []int, i int, j int) {
	s[i], s[j] = s[j], s[i]
}