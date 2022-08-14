package utilz

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Float interface {
	~float32 | ~float64
}

type Integer interface {
	Int | Uint
}

type Number interface {
	Integer | Float
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

func PageOffset[T Integer](pageNum, pageSize T) T {
	return (pageNum - 1) * pageSize
}

func PageCount[T Integer](totalRows, pageSize T) T {
	return CeilDiv(totalRows, pageSize)
}
