package linearsearch

func LinearSearch(arr []string, search string) (int, bool) {
	for i, v := range arr {
		if v == search {
			return i, true
		}
	}
	return 0, false
}
