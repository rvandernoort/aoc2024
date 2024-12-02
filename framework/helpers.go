package framework

func AbsDiffUint(x, y uint64) uint64 {
	if x < y {
		return y - x
	}
	return x - y
}

func Remove(slice []uint64, s int) []uint64 {
	return append(slice[:s], slice[s+1:]...)
}
