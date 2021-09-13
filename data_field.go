package main

import (
	"crypt/lib"
	"fmt"
	"strconv"
)

const INT_SIZE = 4

func GetBoolField(_err error, str string, startAt int64) (err error, field bool, cursor int64) {
	if _err != nil {
		return err, false, cursor
	}
	err = nil
	field = str[startAt:startAt+1] == "1"
	cursor = startAt + 1
	return
}

func ToBoolField(value bool) string {
	if value {
		return "1"
	} else {
		return "0"
	}
}

func GetIntField(_err error, str string, startAt int64) (err error, field int64, cursor int64) {
	if _err != nil {
		return err, 0, cursor
	}
	endAt := startAt + INT_SIZE
	sfield := str[startAt:endAt]
	cursor = endAt
	field, err = strconv.ParseInt("0x"+sfield, 0, 64)
	return
}

func ToIntField(value int64) string {
	hex_value := fmt.Sprintf("%x", value)
	return lib.PadLeft(hex_value, '0', INT_SIZE)
}

func GetStrField(_err error, str string, startAt int64) (err error, field string, cursor int64) {
	if _err != nil {
		return err, "", cursor
	}
	err, fieldLength, cursor := GetIntField(err, str, startAt)
	if err != nil {
		return err, "", cursor
	}
	if len(str) < int(cursor+fieldLength) {
		return fmt.Errorf("length %d out of bounds at cursor %d, slice %s has length %d", fieldLength, cursor, str, len(str)), str, cursor
	}

	field = str[cursor : cursor+fieldLength]
	cursor = cursor + fieldLength
	return
}

func ToStrField(value string) string {
	return ToIntField(int64(len(value))) + value
}
