package utils

func Int32Compare(a *int32, b int) bool {
	if a == nil {
		return false
	}
	return *a == int32(b)
}
