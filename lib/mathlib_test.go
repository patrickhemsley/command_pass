package lib

import (
	"testing"
)

func TestLog(t *testing.T) {
	cases2 := []TTestIntToInt{
		{"TestLog", Log2, 0, 0},
		{"TestLog", Log2, 1, 1},
		{"TestLog", Log2, 2, 2},
		{"TestLog", Log2, 3, 2},
		{"TestLog", Log2, 4, 3},
		{"TestLog", Log2, 5, 3},
	}

	TestIntToInt(t, cases2)
}

func TestNextPowOf2(t *testing.T) {
	cases2 := []TTestIntToInt{
		{"TestLog", NextPowOf2, 3, 4},
		{"TestLog", NextPowOf2, 8, 8},
	}

	TestIntToInt(t, cases2)
}

func TestBits(t *testing.T) {
	cases2 := []TTestIntToInt{
		{"TestBits", Bits, 0, 1},
		{"TestBits", Bits, 1, 2},
		{"TestBits", Bits, 2, 4},
	}

	TestIntToInt(t, cases2)
}
