package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	card, ok := GetItem(slice, index)
	if ok {
		return card, true
	} else {
		return card, false
	}
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	return SetItem(slice, index, value)
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	//left := make([]int, length)
	if length <= 0 {
		return []int{value}
	} else {
		return []int{value}
	}
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	return RemoveItem(slice, index)
}
