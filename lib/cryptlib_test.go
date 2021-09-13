package lib

import (
	"testing"
)

func TestXOR(t *testing.T) {
	cases2 := []TTestStrStrToStrErr{
		{"TestXOR", feistelXOR, "01", "10", "11"},
		{"TestXOR", feistelXOR, "f0", "0f", "ff"},
		{"TestXOR", feistelXOR, "77", "22", "55"},
	}

	TestStrStrToStrErr(t, cases2)
}

func TestFeistel(t *testing.T) {
	//L2, R2 = R1, XOR(L1, F(R1, K1))

	cases2 := []TTestStrStrToStrErr{
		//{"TestFeistelEncryptDecrypt", feistelEncrypt, "0102", "00", "0203"},
		//{"TestFeistelEncryptDecrypt", feistelDecrypt, "0203", "00", "0102"},
		{"TestFeistelEncryptDecrypt", feistelEncrypt, "0102", "00", "0211"},
		{"TestFeistelEncryptDecrypt", feistelDecrypt, "0203", "00", "1302"},
	}

	TestStrStrToStrErr(t, cases2)
}

func TestFeistelEncryptDecrypt(t *testing.T) {
	//L2, R2 = R1, XOR(L1, F(R1, K1))

	cases := []string{"0102", "0203", "ff00012345678934"}
	key := "00ff3399eeaabb22"

	for _, tc := range cases {
		cypher, err := feistelEncrypt(tc, key)
		if err != nil {
			t.Fatalf("feistelEncrypt of %s failed with error %s", tc, err)
			return
		}
		plain, err := feistelDecrypt(cypher, key)
		if err != nil {
			t.Fatalf("feistelDecrypt of %s failed with error %s", cypher, err)
			return
		}
		if tc != plain {
			t.Fatalf("decryption %s of encryption %s differs from original %s", plain, cypher, tc)
			return
		}
	}

}

func TestEncryptDecrypt(t *testing.T) {
	//L2, R2 = R1, XOR(L1, F(R1, K1))

	cases := []string{"Hello", "World", "How are you?"}
	password := "mysecretpassword"

	for _, tc := range cases {
		cypher, err := Encrypt(tc, password)
		if err != nil {
			t.Fatalf("Encrypt of %s failed with error %s", tc, err)
			return
		}
		plain, err := Decrypt(cypher, password)
		if err != nil {
			t.Fatalf("Decrypt of %s failed with error %s", cypher, err)
			return
		}
		if plain != tc {
			//t.Fatalf("decryption %s of encryption %s differs from original %s", plain, cypher, tc)
			t.Fatalf("decryption failed '%s' != '%s', len %d", plain, tc, len(plain))
			return
		}
	}

}

func TestEncryptDecrypt2(t *testing.T) {
	//L2, R2 = R1, XOR(L1, F(R1, K1))

	key := "62958372"

	cypher := "01020304"
	t.Logf("encrypt: %s", cypher)
	for i := 0; i < 20; i++ {
		cypher, _ = feistelEncrypt(cypher, key)
		t.Logf("encrypt -> %s", cypher)
	}
	//t.Fatalf("")

}
