package helpers

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func FindInStringSlice(slice []string, val string) bool {
	for _, n := range slice {
		if val == n {
			return true
		}
	}
	return false
}
