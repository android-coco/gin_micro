package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// 私钥生成
//openssl genrsa -out rsa_private_key.pem 1024
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQCFQCA1mHqQ6LL/MhD2s+gqZew1q6eGDOQpeKjlKj10ndAp8gxI
GzlS1XUVKlYQ5Y/iJlEXUh78kdG8ra9C433JDOVUvVJ7X/I5gB5sknF7uJLytSuS
Qv8G0Y+BK9g8nIctWY97cIb2XW9kbLSsB6h4G9ph9Rm6TlgH2gprhOXWUwIDAQAB
AoGACi90JtSgbcozwHkBvfHTicYfr5yO4hrDJ/5enqHDb9IOUt57HNnj4FaLrBH/
4SvC+0zlfuxajQDScOMv1eOQvjs6v2615dY0nTXYRm6NDehS+VoSK9DZi0BWoE5G
hiSJhckhPCalqHeJga0A5GQtP5q2+hvA2I8gDu7RGVsdsCECQQCFQd24SVUazhxV
k7c1JCKDKK8byq364ulPMaD+SiYgmoTX9stjaBTAKts3TDDE/DV8nc5DVuao29/q
tvwjhcszAkEA//yoIbaQlZroLLBbbqb8LplSIrdjoQk6QrlrT4WIQvY5OMyelICB
8JEaUJdmOmMXkdfwVQmC+2IKGN72dbzIYQJAWmBxn6scrTFcxi2I8+GuBoZxPMgZ
dy6uTae7KLvhX/tsXYxkJOdSK4LlanuiF/d1zy631bP6fEujcezo1K7JQQJAFpWd
58uJmgleroKormyBF0Njobh4S77aqwRc2Vk4ml/K0J4M56Em1aiXn8Cbvk77x1w7
0eTS74bIyUTyjZSoQQJASeDlAmI0o7WQldBFJhncw+RpBDvoFKKoHwIqKR/IjFp4
WYCJK1oqAKLWt0BPOfkDNOkvGtrg3rqV4i+QidP/bw==
-----END RSA PRIVATE KEY-----
`)

// 公钥: 根据私钥生成
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCFQCA1mHqQ6LL/MhD2s+gqZew1
q6eGDOQpeKjlKj10ndAp8gxIGzlS1XUVKlYQ5Y/iJlEXUh78kdG8ra9C433JDOVU
vVJ7X/I5gB5sknF7uJLytSuSQv8G0Y+BK9g8nIctWY97cIb2XW9kbLSsB6h4G9ph
9Rm6TlgH2gprhOXWUwIDAQAB
-----END PUBLIC KEY-----
`)

// 公钥加密
func Encrypt(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 私钥解密
func Decrypt(ciphertext []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

//私钥签名
func SignByPrivate(data []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	//获取私钥
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed)
}

//公钥验证
func SignVer(data []byte, signature []byte) error {
	hashed := sha256.Sum256(data)
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//验证签名
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
}
