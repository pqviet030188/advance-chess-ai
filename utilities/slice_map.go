package utilities

func SliceMap[T, U any](src []T, mapFn func(T) U) []U {
	dest := make([]U, len(src))
	for i := range src {
		dest[i] = mapFn(src[i])
	}
	return dest
}
