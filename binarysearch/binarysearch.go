package binarysearch

func BinarySearch(arr []string, s string) (int, bool) {
	if len(arr) == 0 {
		return 0, false
	}
	start, end := 0, len(arr)-1
	// integer division will round down the result
	mid := (start + end) / 2
	// when start is bigger than end, it means cannot find s string
	for arr[mid] != s && start <= end {
		if s > arr[mid] {
			// increment is to prevent stalling at mid, making the next start closer to the end
			start = mid + 1
		} else {
			// decrement is to prevent stalling at mid, making the end closer to the start
			end = mid - 1
		}
		mid = (start + end) / 2
	}
	if arr[mid] == s {
		return mid, true
	}
	return 0, false

}
