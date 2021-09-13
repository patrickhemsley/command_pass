package lib

func PadPower2(str string) string {
	//pad string to 2^n characters
	return PadRight(str, '0', NextPowOf2(len(str)))
}

func PadRight(str string, pad byte, length int) string {
	charsToAdd := length - len(str)
	for i := 0; i < charsToAdd; i++ {
		str += string(pad)
	}
	return str
}

func PadLeft(str string, pad byte, length int) string {
	charsToAdd := length - len(str)
	for i := 0; i < charsToAdd; i++ {
		str = string(pad) + str
	}
	return str
}

func StrRepeat(str string, size int) string {
	result := ""
	for len(result) < size {
		result += str
	}
	if len(result) > size {
		result = result[0:size]
	}
	return result
}
