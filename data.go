package main

type TItem struct {
	Key   string
	Value string
}

type TData struct {
	Header TOptions
	Items  []TItem
}

//str is not in hex format
func StrToData(str string) (error, TData) {
	var result TData
	var err error

	//File format <header size><header><items size><items>
	cursor := int64(0)
	err, result.Header.version, cursor = GetIntField(err, str, cursor)
	err, result.Header.allowListKeys, cursor = GetBoolField(err, str, cursor)
	err, result.Header.password, cursor = GetStrField(err, str, cursor)
	sitem := str[cursor:]
	err, items := StrToItems(err, sitem)
	result.Items = items
	return err, result
}

func StrToItems(_err error, str string) (err error, items []TItem) {
	if _err != nil {
		return _err, items
	}

	err, itemCount, cursor := GetIntField(err, str, int64(0))
	for i := 0; i < int(itemCount); i++ {
		var item TItem
		err, item, cursor = StrToItem(err, str, cursor)
		items = append(items, item)
	}

	return err, items
}

func StrToItem(_err error, str string, _cursor int64) (err error, item TItem, cursor int64) {
	err, item.Key, cursor = GetStrField(_err, str, _cursor)
	err, item.Value, cursor = GetStrField(_err, str, cursor)
	return
}

func ItemToStr(item TItem) string {
	return ToStrField(item.Key) + ToStrField(item.Value)
}

func ItemsToStr(Items []TItem) string {
	sItemCount := ToIntField(int64(len(Items)))
	result := sItemCount
	for i := 0; i < len(Items); i++ {
		result += ItemToStr(Items[i])
	}
	return result
}

func DataToStr(data TData) string {
	headerStr := OptionsToStr(data.Header)
	itemStr := ItemsToStr(data.Items)
	return headerStr + itemStr
}

func AddItem(data TData, item TItem) TData {
	data.Items = append(data.Items, item)
	return data
}
