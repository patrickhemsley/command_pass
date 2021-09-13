package main

type TFlags = struct {
	allowListKeys bool
}

func strToFlags(str string) TFlags {
	var result TFlags
	result.allowListKeys = str[4:5] == "1"
	return result
}

func flagsToStr(flags TFlags) string {
	var result string
	if flags.allowListKeys {
		result = "00001"
	} else {
		result = "00000"
	}
	return result
}
