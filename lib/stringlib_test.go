package lib

import (
	"testing"
)

func TestPadPower2(t *testing.T) {
	cases2 := []TTestStrToStr{
		{"PadPower2", PadPower2, "12", "12"},
		{"PadPower2", PadPower2, "123", "1230"},
		{"PadPower2", PadPower2, "6f0000004a1c1f09", "6f0000004a1c1f09"},
	}

	TestStrToStr(t, cases2)
}

func TestStrRepeat(t *testing.T) {
	cases2 := []TTestStrIntToStr{
		{"StrRepeat", StrRepeat, "123", 1, "1"},
		{"StrRepeat", StrRepeat, "123", 3, "123"},
		{"StrRepeat", StrRepeat, "123", 7, "1231231"},
	}

	TestStrIntToStr(t, cases2)
}
