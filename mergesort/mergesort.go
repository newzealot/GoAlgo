package mergesort

func merge(arr1 []string, arr2 []string) []string {
	arr := []string{}
	i := 0
	j := 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			arr = append(arr, arr1[i])
			i += 1
		} else {
			arr = append(arr, arr2[j])
			j += 1
		}
	}
	if i < len(arr1) {
		arr = append(arr, arr1[i:]...)
	} else {
		arr = append(arr, arr2[j:]...)
	}
	return arr
}

func MergeSort(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}
