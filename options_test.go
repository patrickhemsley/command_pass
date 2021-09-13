package main

import (
	"crypt/lib"
	"fmt"
	"testing"
)

func TestOptionsToStr(t *testing.T) {
	var options TOptions

	options.version = 1
	options.allowListKeys = true
	options.password = "mypassword"

	actual := ToStrField("mypassword")
	expected := "000amypassword"
	lib.TestAssertBool(t, actual == expected, fmt.Sprintf("TestOptionsToStr: expected %s actual %s", expected, actual))

	actual = OptionsToStr(options)
	expected = "00011000amypassword"
	lib.TestAssertBool(t, actual == expected, fmt.Sprintf("TestOptionsToStr: expected %s actual %s", expected, actual))

	err, options2 := StrToOptions(nil, actual)
	lib.TestAssertErr(t, err, "TestOptionsToStr: StrToOptions")
	actual = options2.password
	expected = options.password
	lib.TestAssertBool(t, actual == expected, fmt.Sprintf("TestOptionsToStr: expected %s actual %s", expected, actual))

}

func TestStrToOptions(t *testing.T) {

	var options TOptions

	options.version = 1
	options.allowListKeys = true
	options.password = "mypassword"
	str := OptionsToStr(options)

	var err error
	var result TOptions

	cursor := int64(0)
	err, result.version, cursor = GetIntField(err, str, cursor)
	lib.TestAssertErr(t, err, "TestOptionsToStr: GetIntField")
	err, result.allowListKeys, cursor = GetBoolField(err, str, cursor)
	lib.TestAssertErr(t, err, "TestOptionsToStr: GetBoolField")
	err, result.password, cursor = GetStrField(err, str, cursor) //TODO use RUNE instead
	lib.TestAssertErr(t, err, "TestOptionsToStr: GetStrField")

	expectedInt := options.version
	actualInt := result.version
	lib.TestAssertBool(t, actualInt == expectedInt, fmt.Sprintf("TestOptionsToStr: options %s expected %d actual %d", str, expectedInt, actualInt))

	expectedBool := options.allowListKeys
	actualBool := result.allowListKeys
	lib.TestAssertBool(t, actualBool == expectedBool, fmt.Sprintf("TestOptionsToStr: options %s expected %d actual %d", str, expectedInt, actualInt))

	expected := options.password
	actual := result.password
	lib.TestAssertBool(t, actual == expected, fmt.Sprintf("TestOptionsToStr: options %s expected %s actual %s", str, expected, actual))
}

func TestGetIntField(t *testing.T) {
	str := "0001"
	startAt := int64(0)
	var err error
	err, version, _ := GetIntField(err, str, startAt)
	lib.TestAssertBool(t, version == 1, fmt.Sprintf("TestIntField: expected %d actual %d", 1, version))
	lib.TestAssertErr(t, err, "TestIntField")

}

func TestGetStrField(t *testing.T) {
	str := "000amypassword"
	startAt := int64(0)
	var err error
	err, actual, _ := GetStrField(err, str, startAt)
	expected := "mypassword"

	lib.TestAssertBool(t, actual == expected, fmt.Sprintf("TestIntField: expected %s actual %s", expected, actual))
	lib.TestAssertErr(t, err, "TestGetStrField")

}
