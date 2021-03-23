package quicksort

func Pivot(arr []string, start int, end int) (swapIndex int) {
	pivot := arr[start]
	swapIndex = start
	for i := start + 1; i <= end; i++ {
		if arr[i] < pivot {
			swapIndex += 1
			arr[i], arr[swapIndex] = arr[swapIndex], arr[i]
		}
	}
	arr[start], arr[swapIndex] = arr[swapIndex], arr[start]
	return swapIndex
}

func QuickSort(arr []string, left int, right int) []string {
	if left < right {
		pivotIndex := Pivot(arr, left, right)
		QuickSort(arr, left, pivotIndex-1)
		QuickSort(arr, pivotIndex+1, right)
	}
	return arr
}
