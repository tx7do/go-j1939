package math

// Min 求int64型的最小值
func Min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

// Max 求int64型的最大值
func Max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
