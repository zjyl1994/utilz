package utilz

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func CeilDiv[T Integer](x, y T) T {
	return (x + y - 1) / y
}

func Min[T Number](nums ...T) T {
	min := nums[0]

	for _, v := range nums {
		if min > v {
			min = v
		}
	}

	return min
}

func Max[T Number](nums ...T) T {
	max := nums[0]

	for _, v := range nums {
		if max < v {
			max = v
		}
	}

	return max
}
