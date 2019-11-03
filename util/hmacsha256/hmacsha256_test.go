package hmacsha256

import (
	"crypto"
	"fmt"
	"testing"
)

func TestHmacSha2562(t *testing.T) {
	fmt.Println(HmacSha256("hello","111"))
}

func TestHmacEncrypt(t *testing.T) {
	fmt.Println(HmacEncrypt([]byte("hello"),[]byte("111"),crypto.SHA256))
}

func TestHmacEncryptToBase64(t *testing.T) {
	fmt.Println(HmacEncryptToBase64([]byte("hello"),[]byte("111"),crypto.SHA256))
}
