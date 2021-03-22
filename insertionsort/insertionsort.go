package insertionsort

func InsertionSort(arr []string) []string {
	if len(arr) < 2 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		// Have to store arr[i] in currentValue as values in arr[i] will keep changing
		currentValue := arr[i]
		// j have to be stored outside of "for" scope in order to execute "arr[j+1] = currentValue"
		j := i - 1
		for ; j >= 0 && arr[j] > currentValue; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = currentValue
	}
	return arr
}
