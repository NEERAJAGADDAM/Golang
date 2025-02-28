package remove

func RemoveKeyFromSlice(slice []int, key int) []int {
	for i, k := range slice {
		if k == key {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
