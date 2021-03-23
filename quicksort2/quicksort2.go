package quicksort2

func partition(arr []string, leftIndex int, rightIndex int) int {
	pivotValue := arr[(leftIndex+rightIndex)/2]
	for leftIndex < rightIndex {
		for arr[leftIndex] < pivotValue {
			leftIndex += 1
		}
		for arr[rightIndex] > pivotValue {
			rightIndex -= 1
		}
		if leftIndex < rightIndex {
			arr[leftIndex], arr[rightIndex] = arr[rightIndex], arr[leftIndex]
		}
	}
	return leftIndex
}

func QuickSort(arr []string, leftIndex int, rightIndex int) []string {
	if leftIndex < rightIndex {
		pivotIndex := partition(arr, leftIndex, rightIndex)
		QuickSort(arr, leftIndex, pivotIndex-1)
		QuickSort(arr, pivotIndex+1, rightIndex)
	}
	return arr
}
