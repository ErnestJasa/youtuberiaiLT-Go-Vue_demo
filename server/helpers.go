package main

func Filter[T any](items []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func All[T any](items []T, predicate func(T) bool) bool {
	for _, item := range items {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func Any[T any](items []T, predicate func(T) bool) bool {
	for _, item := range items {
		if predicate(item) {
			return true
		}
	}
	return false
}

func RemoveEmptyStrings(strings []string) []string {
	var result = strings
	for i, string := range strings {
		if string == "" {
			result = append(result[:i], result[i+1:]...)
		}
	}
	return result
}
