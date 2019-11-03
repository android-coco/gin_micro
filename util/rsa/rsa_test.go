package rsa

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestRsaDecrypt(t *testing.T) {
	mw := []byte("ceshi01ceshi01ceshi01ceshi01ceshi01ceshi01ceshi01ceshi01ceshi01ceshi01")
	fmt.Println(len(mw))
	data, err := Encrypt(mw)
	fmt.Println(base64.StdEncoding.EncodeToString(data),err)
	//bytes, _ := base64.StdEncoding.DecodeString("uhPDW8jD4q5cbVIti7c540eR1ZXXLWR3UTcxVzrC1R+1rd2+9EPTz8J+N/uphWkR5x2Yq94bnAGOi3EK7icd2BgAT67gi7mccm2kbplYTND5CMSJ4QmoD2hcl8sOFVNPzko8A/0rMT1J+zerv9NXEfXfujwALmK+BhSZVbvN7OY=")
	origData, err := Decrypt(data)
	fmt.Println(string(origData),err)

	private, err := SignByPrivate([]byte("fad"))
	fmt.Println("1111:",err)
	err = SignVer([]byte("fad"), private)
	fmt.Println("2222:",err)
}
