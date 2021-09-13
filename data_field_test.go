package main

import (
	"crypt/lib"
	"testing"
)

func TestToField(t *testing.T) {

	lib.TestAssertBool(t, ToIntField(3) == "0003", "ToIntField: "+ToIntField(3))
	lib.TestAssertBool(t, ToStrField("abc") == "0003abc", "ToStrField: "+ToStrField("abc"))
	lib.TestAssertBool(t, ToStrField("mypassword") == "000amypassword", "ToStrField: "+ToStrField("mypassword"))

}
