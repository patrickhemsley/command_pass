package main

const MAX_PASSWORD_LENGTH = 64
const OPTIONS_LENGTH = MAX_PASSWORD_LENGTH + 3 + 5

type TOptions struct {
	version       int64
	allowListKeys bool
	password      string
}

func StrToOptions(_err error, str string) (err error, result TOptions) {

	if _err != nil {
		return _err, result
	}
	cursor := int64(0)
	err = nil

	err, result.version, cursor = GetIntField(err, str, cursor)
	err, result.allowListKeys, cursor = GetBoolField(err, str, cursor)
	err, result.password, cursor = GetStrField(err, str, cursor) //TODO use RUNE instead
	return err, result
}

func OptionsToStr(options TOptions) string {

	return ToIntField(options.version) + ToBoolField(options.allowListKeys) + ToStrField(options.password)
}
