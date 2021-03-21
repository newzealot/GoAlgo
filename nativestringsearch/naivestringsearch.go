package nativestringsearch

func NativeStringSearch(arr string, s string) (counter int) {
	if len(s) > len(arr) {
		return 0
	}
	i := 0
	for i <= len(arr)-len(s) {
		if arr[i:i+len(s)] == s {
			counter += 1
		}
		i += 1
	}
	return counter
}
