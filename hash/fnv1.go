package hash

import "hash/fnv"

func Fnv1(data []byte) uint64 {
	h := fnv.New64()
	h.Write(data)
	return h.Sum64()
}

func Fnv1String(s string) uint64 {
	return Fnv1([]byte(s))
}
