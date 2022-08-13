package utilz

func Contains[T comparable](slice []T, element T) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}

	return false
}

func Filter[V any](slice []V, fn func(V) bool) []V {
	result := []V{}

	for _, item := range slice {
		if fn(item) {
			result = append(result, item)
		}
	}

	return result
}

func Map[T any, R any](slice []T, fn func(T) R) []R {
	result := make([]R, len(slice))

	for i, item := range slice {
		result[i] = fn(item)
	}

	return result
}

func Uniq[T comparable](slice []T) []T {
	result := make([]T, 0, len(slice))
	flag := make(map[T]struct{}, len(slice))

	for _, item := range slice {
		if _, ok := flag[item]; ok {
			continue
		}

		flag[item] = struct{}{}
		result = append(result, item)
	}

	return result
}

func Keys[K comparable, V any](m map[K]V) []K {
	result := make([]K, 0, len(m))

	for k := range m {
		result = append(result, k)
	}

	return result
}

func Values[K comparable, V any](m map[K]V) []V {
	result := make([]V, 0, len(m))

	for _, v := range m {
		result = append(result, v)
	}

	return result
}
