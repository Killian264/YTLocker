package helpers

func IsKeyInArray(array []uint64, key uint64) bool {
	for _, num := range array {
		if num == key {
			return true
		}
	}

	return false
}
