package main

import (
	"fmt"
)

//TAction use case
type TAction int

//Valid actions
const (
	NoAction TAction = iota
	InitFile
	AddKeyValue
	DeleteKeyValue
	CopyValue
	ListKeys
	CheckFile
)

//TParam program parameters
type TParam struct {
	action        TAction
	passwordFile  string
	key           string
	value         string
	allowListKeys bool
	password      string
}

func getParam(args []string) (TParam, error) {
	var p TParam
	const DEFAULT_FILE = "datafile"

	if len(args) == 1 {
		return p, fmt.Errorf("parameter expected")
	}

	p.passwordFile = DEFAULT_FILE
	p.allowListKeys = true

	if len(args) > 5 {
		return p, fmt.Errorf("too many parameters")
	}
	for i := 1; i < len(args); i++ {
		param2 := ""
		param3 := ""
		param4 := ""
		if i < len(args)-1 {
			param2 = args[i+1]
		}
		if i < len(args)-2 {
			param3 = args[i+2]
		}
		if i < len(args)-3 {
			param4 = args[i+3]
		}
		flag1 := ""
		nextArgIsParam := (len(param2) == 2) && (param2[1:2] == "-")
		nextArgIsFlag := !nextArgIsParam && (len(param2) > 2) && (param2[1:2] == "-") && (param2[2:3] != "-")
		if nextArgIsParam {
			param2 = ""
			param3 = ""
			param4 = ""
		} else if nextArgIsFlag {
			flag1 = param2
			param2 = param3
			param3 = param4
			param4 = ""
		}

		if args[i] == "--help" || args[i] == "-h" {
			p.action = NoAction
			fmt.Println("crypt is a program to store a collection of encrypted keys and values (e.g. passwords)")
			fmt.Println("")
			fmt.Println("Note:")
			fmt.Println(" It is possbile though not recommended to pass the password in via the command line with the -p option.")
			fmt.Println("You will be prompted for the password if -p option is not used ")
			fmt.Println("Parameters enclosed in [] are optional")
			fmt.Println("")
			fmt.Println("usage:")
			fmt.Println("crypt [-p password] [-f file path] -i [-nolist]  Initialise password file <file path>.")
			fmt.Println("                                         if option <nolist> is specified then -k option will ")
			fmt.Println("                                         be unavailable ")
			fmt.Println("                                         e.g.: crypt -i mysecrets.txt")
			fmt.Println("crypt  [-p password] [-f file path] -a <key> \"<value>\"     Add a key-value pair to file <file path> ")
			fmt.Println("                                         e.g.: crypt -a mysecrets.txt Bank1 \"My secret password\"")
			fmt.Println("crypt [-p password] [-f file path] -d <key>               Remove a key-value pair from file <file path> ")
			fmt.Println("                                         e.g.: crypt -d mysecrets.txt Bank1")
			fmt.Println("crypt [-p password] [-f file path] -r <key>               Copy to the clipboard the value associated with the key <key> in file <file path> ")
			fmt.Println("                                         You will be prompted for the password")
			fmt.Println("                                         e.g.: crypt -r mysecrets.txt Bank1")
			fmt.Println("crypt [-f file path] -k                    List keys in password file <file path>.")
			fmt.Println("                                         If file was initiliased with the -nolist option")
			fmt.Println("                                         then this options will be unavailable")
			fmt.Println("crypt [-p password] [-f file path] -t                    Check password file <file path>.")

		} else if args[i] == "-i" {
			p.action = InitFile
			if flag1 == "-nolist" {
				p.allowListKeys = false
				i++
			}
		} else if args[i] == "-k" {
			p.action = ListKeys
		} else if args[i] == "-p" {
			p.password = param2
			i++
		} else if args[i] == "-f" {
			p.passwordFile = param2
			i++
		} else if args[i] == "-t" {
			p.action = CheckFile
		} else if args[i] == "-a" {
			p.action = AddKeyValue
			if param3 == "" {
				return p, fmt.Errorf("parameter expected")
			} else {
				p.key = param2
				p.value = param3
				i = i + 2
			}
		} else if args[i] == "-d" {
			p.action = DeleteKeyValue
			if param2 == "" {
				return p, fmt.Errorf("parameter expected")
			} else if param3 == "" {
				p.key = param2
				i = i + 1
			} else {
				return p, fmt.Errorf("unexpected extra parameter")
			}
		} else if args[i] == "-r" {
			p.action = CopyValue
			if param2 == "" {
				return p, fmt.Errorf("parameter expected")
			} else if param3 == "" {
				p.key = param2
				i = i + 1
			} else {
				return p, fmt.Errorf("unexpected extra parameter")
			}
		} else {
			return p, fmt.Errorf("'" + args[i] + "' not implemented")
		}
	}
	return p, nil
}
