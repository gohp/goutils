package sort

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
* @Author: Jam Wong
* @Date: 2020/6/22
 */

var randomList = []int{2, 7, 19, 45, 3, 18, 5, 9, 35, 26}
var expectList = []int{2, 3, 5, 7, 9, 18, 19, 26, 35, 45}

func TestBubbleSort(t *testing.T) {
	Convey("Test Bubble Sort", t, func() {
		So(BubbleSort(randomList), ShouldResemble, expectList)
	})
}

func TestSelectSort(t *testing.T) {
	Convey("Test Select Sort", t, func() {
		So(SelectSort(randomList), ShouldResemble, expectList)
	})
}

func TestInsertSort(t *testing.T) {
	Convey("Test Insert Sort", t, func() {
		So(InsertSort(randomList), ShouldResemble, expectList)
	})
}

func TestQuickSort(t *testing.T) {
	Convey("Test Quick Sort", t, func() {
		So(QuickSort(randomList), ShouldResemble, expectList)
	})
}

func TestMergingSort(t *testing.T) {
	Convey("Test Merging Sort", t, func() {
		So(MergingSort(randomList), ShouldResemble, expectList)
	})
}

func TestHeapSort(t *testing.T) {
	Convey("Test Heap Sort", t, func() {
		So(HeapSort(randomList), ShouldResemble, expectList)
	})
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(randomList)
	}
}

func BenchmarkSelectSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectSort(randomList)
	}
}

func BenchmarkInsertSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertSort(randomList)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(randomList)
	}
}

func BenchmarkMergingSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergingSort(randomList)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HeapSort(randomList)
	}
}
