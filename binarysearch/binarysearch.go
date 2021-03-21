package binarysearch

func BinarySearch(arr []string, search string) (int, bool) {
	lIndex := 0
	rIndex := len(arr) - 1
	for lIndex <= rIndex {
		mIndex := (rIndex + lIndex) / 2
		if arr[mIndex] == search {
			return mIndex, true
		}
		if arr[mIndex] > search {
			rIndex = mIndex - 1
		}
		if arr[mIndex] < search {
			lIndex = mIndex + 1
		}
	}
	return 0, false
}
