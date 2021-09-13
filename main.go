package main

import (
	"fmt"
	"os"

	"github.com/howeyc/gopass"
)

//PromptPassword prompts user for a password
func PromptPassword() string {
	password := ""
	for password == "" {
		fmt.Print("password: ")
		//bytePassword, _ := terminal.ReadPassword(0)
		bytePassword, err := gopass.GetPasswd()
		if err != nil {
			panic(err)
		}
		password = string(bytePassword)
		fmt.Print("\n")
	}
	return password
}

func ValidPassword(password string) bool {
	result := len(password) > 7
	if !result {
		fmt.Println("password must be at least 8 characters")
	}
	return result
}

func main() {

	var options TOptions
	options.version = 1
	param, err := getParam(os.Args)
	if err != nil {
		fmt.Println("invalid parameters:", err)
		return
	}
	options.password = param.password

	//prompted for the password to use
	switch param.action {
	case InitFile, AddKeyValue, DeleteKeyValue,
		CopyValue, ListKeys, CheckFile:
		if options.password == "" {
			options.password = PromptPassword()
		}
		if !ValidPassword(options.password) {
			return
		}
	}

	switch param.action {
	case InitFile:
		options.allowListKeys = param.allowListKeys
		err = InitialiseFile(param.passwordFile, options)
		if err != nil {
			s := fmt.Sprintf("error initialising password file %s: %s", param.passwordFile, err)
			fmt.Println(s)
			return
		}

	case CheckFile:
		data, err := LoadFile(param.passwordFile, options)
		if err != nil {
			s := fmt.Sprintf("error checking file %s: %s", param.passwordFile, err)
			fmt.Println(s)
			return
		}
		fmt.Println(fmt.Sprintf("debug read: header %s\n", OptionsToStr(data.Header)))

	case ListKeys:

	case AddKeyValue:

	case DeleteKeyValue:

	case CopyValue:

	}

	fmt.Println(param)

}
