package main

import (
	"crypt/lib"
	"errors"
	"fmt"
	"os"
)

//InitFile initialises password file <path> with password <password>
func InitialiseFile(path string, options TOptions) error {
	if _, err := os.Stat(path); err == nil {
		return errors.New(fmt.Sprintf("cannot initialise file %s since it already exists", path))
	}
	var data TData
	data.Header = options

	header := OptionsToStr(options)
	xheader, err := lib.Encrypt(header, options.password)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("debug: write: %s -> %s\n", header, xheader))
	err = lib.WriteToTextFile(path, xheader)
	if err != nil {
		return err
	}

	return nil
}

//Loads password file
func LoadFile(path string, options TOptions) (TData, error) {
	var result TData

	contentString, err := lib.ReadFromTextFile(path)
	if err != nil {
		return result, err
	}
	decrypted, err := lib.Decrypt(contentString, options.password)
	if err != nil {
		return result, err
	}
	err, result = StrToData(decrypted)
	return result, nil
}
