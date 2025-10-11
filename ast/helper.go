package ast

func Cast[T any](stm Statement) T {
	var defaultVal T
	if v, ok := stm.(T); ok {
		return v
	}

	return defaultVal
}
