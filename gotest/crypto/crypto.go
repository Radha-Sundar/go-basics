package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

var IV = "facd343d4d6a3378e176504ea94f116a"

func EncryptAes128Cbc(key string, plaintext string) string {
	iv, _ := hex.DecodeString(string(IV)) //nolint:errcheck //simplify
	b := addPkcs7Padding([]byte(plaintext))

	block, _ := aes.NewCipher(_md5(key)) //nolint:errcheck //simplify
	ciphertext := make([]byte, len(b))

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, b)

	return hex.EncodeToString(ciphertext)
}

func DecryptAes128Cbc(key string, ct string) string {
	iv, _ := hex.DecodeString(IV)         //nolint:errcheck //simplify
	ciphertext, _ := hex.DecodeString(ct) //nolint:errcheck //simplify

	block, _ := aes.NewCipher(_md5(key)) //nolint:errcheck //simplify
	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(ciphertext, ciphertext)

	return string(removePkcs7Padding(ciphertext))
}

func _md5(key string) []byte {
	h := md5.New() //nolint:gosec //md5 is ok
	_, _ = io.WriteString(h, key)

	return h.Sum(nil)
}

func addPkcs7Padding(plaintext []byte) []byte {
	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(plaintext, padText...)
}

func removePkcs7Padding(ciphertext []byte) []byte {
	padSize := int(ciphertext[len(ciphertext)-1])
	return ciphertext[:len(ciphertext)-padSize]
}

func Sha256(data, secret []byte) string {
	h := hmac.New(sha256.New, secret)
	h.Write(data)

	return hex.EncodeToString(h.Sum(nil))
}
