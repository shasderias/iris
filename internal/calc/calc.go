package calc

func Abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

func ModClamp(len, idx int) int {
	return ((idx % len) + len) % len
}

func IdxWrap[T any](slice []T, index int) T {
	return slice[ModClamp(len(slice), index)]
}
