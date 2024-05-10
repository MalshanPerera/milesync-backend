package utils

func TransformSlice[T any, U any](data []T, mapper func(T) U) []U {
	result := make([]U, len(data))
	for i, element := range data {
		result[i] = mapper(element)
	}
	return result
}
