package lib

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/bits"
	"strconv"
	"strings"
)

const PASSES = 4

func Encrypt(str string, password string) (string, error) {
	inHex := hex.EncodeToString([]byte(str)) //convert to hex
	inHex = PadPower2(inHex)
	passHex := hex.EncodeToString([]byte(password)) //convert to hex
	for i := 0; i < PASSES; i++ {
		s, err := feistelEncrypt(inHex, passHex)
		if err != nil {
			return "", err
		}
		inHex = s
	}
	return inHex, nil
}

func Decrypt(str string, password string) (string, error) {
	inHex := str
	println(fmt.Sprintf("Decrypt: str=%s, inHex=%s", str, inHex))
	inHex = PadPower2(inHex)
	passHex := hex.EncodeToString([]byte(password)) //convert to hex
	println(fmt.Sprintf("Decrypt: str=%s, password=%s, passHex=%s, inHex=%s", str, password, passHex, inHex))
	for i := 0; i < PASSES; i++ {
		s, err := feistelDecrypt(inHex, passHex)
		if err != nil {
			return "", err
		}
		inHex = s
		println(fmt.Sprintf("Decrypt: inHex=%s", inHex))
	}
	chars, err := hex.DecodeString(inHex) //convert to string
	result := strings.TrimRight(string(chars), "\x00")
	return result, err
}

func feistelFn(hexstr1 string, hexstr2 string) (string, error) {
	result := ""
	if len(hexstr1) != len(hexstr2) {
		return result, errors.New("feistelXOR: expected arguments of same length")
	}
	for i := 0; i < len(hexstr1); i += 2 {
		hex := hexstr1[i : i+2]
		var u1, u2 uint8
		b1, err := strconv.ParseInt(hex, 16, 64)
		u1 = uint8(b1)
		if err != nil {
			return "", err
		}
		hex = hexstr2[i : i+2]
		b2, err := strconv.ParseInt(hex, 16, 64)
		u2 = uint8(b2)
		if err != nil {
			return "", err
		}

		u3 := bits.RotateLeft8(u1, 3) ^ bits.RotateLeft8(u2, 3)
		result += fmt.Sprintf("%02x", u3)
	}
	return result, nil
}

func feistelEncrypt(str string, key string) (string, error) {
	//L2, R2 = R1, XOR(L1, F(R1, K1))
	if str == "" {
		return "", nil
	}
	l := len(str)
	p := NextPowOf2(l)
	if p != l {
		return "", errors.New(fmt.Sprintf("FeistelEncrypt: argument of incorrect length, expected %d not %d", l, p))
	}
	L1 := str[0 : l/2]
	R1 := str[l/2 : l]
	L2 := R1
	K1 := StrRepeat(key, l/2)
	println(fmt.Sprintf("FeistelEncrypt: str=%s, key=%s, K1=%s", str, key, K1))
	F, err := feistelFn(R1, K1)
	if err != nil {
		return "", err
	}
	R2, err := feistelXOR(L1, F)
	if err != nil {
		return "", err
	}
	return L2 + R2, nil
}

func feistelDecrypt(str string, key string) (string, error) {
	//L1, R1 = XOR(R2, F(L2, K1)), L2

	if str == "" {
		return "", nil
	}
	l := len(str)
	p := NextPowOf2(l)
	if p != l {
		return "", errors.New(fmt.Sprintf("FeistelEncrypt: argument of incorrect length, expected %d not %d", l, p))
	}
	L2 := str[0 : l/2]
	R2 := str[l/2 : l]
	R1 := L2
	K1 := StrRepeat(key, l/2)
	println(fmt.Sprintf("FeistelDecrypt: str=%s, key=%s, K1=%s", str, key, K1))
	F, err := feistelFn(L2, K1)
	if err != nil {
		return "", err
	}
	L1, err := feistelXOR(R2, F)
	if err != nil {
		return "", err
	}
	return L1 + R1, nil
}

func feistelXOR(hexstr1 string, hexstr2 string) (string, error) {
	result := ""
	if len(hexstr1) != len(hexstr2) {
		return result, errors.New("feistelXOR: expected arguments of same length")
	}
	for i := 0; i < len(hexstr1); i += 2 {
		hex := hexstr1[i : i+2]
		b1, err := strconv.ParseInt(hex, 16, 64)
		if err != nil {
			return "", err
		}
		hex = hexstr2[i : i+2]
		b2, err := strconv.ParseInt(hex, 16, 64)
		if err != nil {
			return "", err
		}
		b3 := b1 ^ b2
		result += fmt.Sprintf("%02x", b3)
	}
	//fmt.Println(fmt.Sprintf("feistelXOR(%s, %s) = %s", hexstr1, hexstr2, result))
	return result, nil
}
