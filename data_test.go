package main

import (
	"crypt/lib"
	"testing"
)

/*
func TestDataToStr(t *testing.T) {
	var data TData
	var item TItem

	data.Header.version = 1
	data.Header.allowListKeys = true
	data.Header.password = "mypassword"
	item.Key = "mykey"
	item.Value = "myvalue"
	data = AddItem(data, item)
	strdata := DataToStr(data)
	expected := "00011000amypassword000005mykey000000000007myvalue"
	lib.TestAssertExpectedStr(t, nil, expected, strdata, "TestDataToStr")
	t.Logf("TestDataToStr: DataToStr(data)=%s", strdata)
	err, data2 := StrToData(strdata)
	lib.TestAssertErr(t, err, "TestDataToStr: StrToData")
	lib.TestAssertBool(t, data.Header.password == data2.Header.password, "TestTData2: header password")

	lib.TestAssertBool(t, len(data.Items) == 1, "TestDataToStr: data item size")
	lib.TestAssertBool(t, len(data2.Items) == 1, fmt.Sprintf("TestDataToStr: data2 item size = %d", len(data2.Items)))
	lib.TestAssertBool(t, data.Items[1].Value == "myvalue", "TestDataToStr")

}
*/

func TestAddItem(t *testing.T) {
	var data TData
	var item TItem
	item.Key = "key1"
	item.Value = "value1"
	data = AddItem(data, item)

	lib.TestAssertBool(t, len(data.Items) == 1, "TestTData1: 1")
	item.Value = "value2"
	data = AddItem(data, item)
	lib.TestAssertBool(t, data.Items[1].Value == "value2", "TestTData1: 2")
}

//str is not in hex format
func TestStrToData(t *testing.T) {

	str := "0001" + "1" + "000a" + "mypassword" + "0001" + "0005" + "mykey" + "0007" + "myvalue"
	t.Logf("TestStrToData: str=%s", str)

	var header TOptions
	cursor := int64(0)
	var err error
	err, header.version, cursor = GetIntField(err, str, cursor)
	err, header.allowListKeys, cursor = GetBoolField(err, str, cursor)
	err, header.password, cursor = GetStrField(err, str, cursor)
	lib.TestAssertExpectedStr(t, "mypassword", header.password, "TestStrToData")
	sitem := str[cursor:]
	t.Logf("sitems=" + sitem)
	err, items := StrToItems(err, sitem)
	t.Logf("item count = %d", len(items))
	lib.TestAssertBool(t, len(items) == 1, "TestStrToData")

	err, itemCount, cursor := GetIntField(err, str, cursor)
	lib.TestAssertExpectedInt(t, int(itemCount), 1, "TestStrToData")
	err, key, cursor := GetStrField(err, str, cursor)
	lib.TestAssertExpectedStr(t, key, "mykey", "TestStrToData")
	err, value, cursor := GetStrField(err, str, cursor)
	lib.TestAssertExpectedStr(t, value, "myvalue", "TestStrToData")

}

func TestStrToItems(t *testing.T) {
	str := "0001" + "0004" + "key1" + "0006" + "value1"

	var items []TItem
	var err error
	err, itemCount, cursor := GetIntField(err, str, int64(0))
	for i := 0; i < int(itemCount); i++ {
		var item TItem
		err, item, cursor = StrToItem(err, str, cursor)
		items = append(items, item)
	}
	lib.TestAssertBool(t, len(items) == 1, "TestStrToItems")
	lib.TestAssertExpectedStr(t, "key1", items[0].Key, "TestStrToItems")
	lib.TestAssertExpectedStr(t, "value1", items[0].Value, "TestStrToItems")

}

func TestItemsToStr(t *testing.T) {
	var Items []TItem
	var item TItem
	item.Key = "key1"
	item.Value = "value1"
	Items = append(Items, item)

	sItemCount := ToIntField(int64(len(Items)))
	lib.TestAssertExpectedStr(t, "0001", sItemCount, "TestItemsToStr")
	t.Logf("sItemCount=%s", sItemCount)
	result := sItemCount
	for i := 0; i < len(Items); i++ {
		result += ItemToStr(Items[i])
	}
	expected := "0001" + "0004" + "key1" + "0006" + "value1"
	lib.TestAssertExpectedStr(t, expected, result, "TestItemsToStr")
}
