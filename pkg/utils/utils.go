package utils

func ConvertToMapByKey[K comparable, A any](slice []A, key func(A) K) map[K]A {
	result := make(map[K]A)
	for _, element := range slice {
		result[key(element)] = element
	}
	return result
}
