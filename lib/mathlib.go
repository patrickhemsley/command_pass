package lib

func Bits(n int) int {
	k := 1
	for i := 0; i < n; i++ {
		k = k << 1
	}
	return k
}

func NextPowOf2(n int) int {
	k := 1
	for k < n {
		k = k << 1
	}
	return k
}

func Log2(n int) int {
	k := 0
	for m := n; m > 0; m = m >> 1 {
		k++
	}
	return k
}
