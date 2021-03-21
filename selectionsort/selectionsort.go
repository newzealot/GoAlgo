package selectionsort

import "fmt"

func SelectionSort(arr []string) []string {
	minIndex := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := minIndex + 1; j < len(arr); j++ {
			fmt.Println(minIndex, arr[minIndex], arr[j], arr[j] < arr[minIndex])
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
		minIndex = i + 1
		fmt.Println(arr)
	}
	return arr
}
