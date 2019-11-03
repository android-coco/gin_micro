package hmacsha256

import (
	"crypto"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func HmacSha256(value, secret string) string {

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(value))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}

func HmacEncrypt(origData, key []byte, hash crypto.Hash) string {
	mac := hmac.New(hash.New, key)
	mac.Write(origData)
	return hex.EncodeToString(mac.Sum(nil))
}

func HmacEncryptToBase64(origData, key []byte, hash crypto.Hash) string {
	mac := hmac.New(hash.New, key)
	mac.Write(origData)
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
