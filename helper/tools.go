package helper

func CondValue(cond bool, isTrue, isFalse interface{}) interface{} {
	if cond {
		return isTrue
	}
	return isFalse
}

func CondFunc(cond bool, isTrue, isFalse func() interface{}) interface{} {
	if cond {
		return isTrue()
	}
	return isFalse()
}

// If is a helper function that returns one of two elements based on a condition.
func TernaryIf[T any](condition bool, trueVal T, falseVal T) T {
	if condition {
		return trueVal
	}

	return falseVal
}

func Ptr[T any](v T) *T {
	return &v
}