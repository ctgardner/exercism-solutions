package binarysearch

func SearchInts(list []int, key int) int {
	i, j := 0, len(list)
	for i < j {
		k := i + (j-i)/2
		if key == list[k] {
			return k
		}

		if key < list[k] {
			j = k
		} else {
			i = k + 1
		}
	}
	return -1
}
