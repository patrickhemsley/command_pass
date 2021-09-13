package lib

import (
	"fmt"
	"testing"
)

type TTestIntToInt = struct {
	description string
	function    func(int) int
	argument    int
	expected    int
}

type TTestStrToStr = struct {
	description string
	function    func(string) string
	argument    string
	expected    string
}

type TTestStrStrToStrErr = struct {
	description string
	function    func(str1 string, str2 string) (string, error)
	arg1        string
	arg2        string
	expected    string
}

type TTestStrIntToStr = struct {
	description string
	function    func(str string, size int) string
	arg1        string
	arg2        int
	expected    string
}

func TestIntToInt(t *testing.T, tests []TTestIntToInt) {
	for _, tc := range tests {
		actual := tc.function(tc.argument)
		if actual != tc.expected {
			t.Fatalf("%s: expected: %d got: %d for argument: %d",
				tc.description, tc.expected, actual, tc.argument)
		}
	}
}

func TestStrToStr(t *testing.T, tests []TTestStrToStr) {
	for _, tc := range tests {
		actual := tc.function(tc.argument)
		if actual != tc.expected {
			t.Fatalf("%s: expected: %s got: %s for argument: %s",
				tc.description, tc.expected, actual, tc.argument)
		}
	}
}

func TestStrStrToStrErr(t *testing.T, tests []TTestStrStrToStrErr) {
	for _, tc := range tests {
		actual, err := tc.function(tc.arg1, tc.arg2)
		if err != nil {
			t.Fatalf("%s: error %s for arguments: %s, %s",
				tc.description, err, tc.arg1, tc.arg2)
		}
		if actual != tc.expected {
			t.Fatalf("%s: expected: %s got: %s for arguments: %s, %s",
				tc.description, tc.expected, actual, tc.arg1, tc.arg2)
		}
	}
}

func TestStrIntToStr(t *testing.T, tests []TTestStrIntToStr) {
	for _, tc := range tests {
		actual := tc.function(tc.arg1, tc.arg2)
		if actual != tc.expected {
			t.Fatalf("%s: expected: %s got: %s for arguments: %s, %d",
				tc.description, tc.expected, actual, tc.arg1, tc.arg2)
		}
	}
}

func TestAssertBool(t *testing.T, condition bool, message string) {
	if !condition {
		t.Fatal(message)
	}
}

func TestAssertErr(t *testing.T, err error, message string) {
	if err != nil {
		if message == "" {
			t.Fatal(err.Error())
		} else {
			t.Fatal(message + ": " + err.Error())
		}
	}
}

func TestAssertExpectedStr(t *testing.T, expected string, actual string, context string) {
	if actual != expected {
		var prefix string
		if context == "" {
			prefix = ""
		} else {
			prefix = context + ": "
		}
		t.Fatal(prefix + fmt.Sprintf("expected %s not %s", expected, actual))
	}
}

func TestAssertExpectedInt(t *testing.T, expected int, actual int, context string) {
	if actual != expected {
		var prefix string
		if context == "" {
			prefix = ""
		} else {
			prefix = context + ": "
		}
		t.Fatal(prefix + fmt.Sprintf("expected %d not %d", expected, actual))
	}
}
