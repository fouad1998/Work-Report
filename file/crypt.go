package file

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

var CRYPT_KEY = []byte("YtRltyeCYjhAdVx8kenzSWAjYnvxFoDx")

func encrypt(source []byte) []byte {
	block, _ := aes.NewCipher(CRYPT_KEY)

	cipherText := make([]byte, aes.BlockSize+len(source))
	iv := cipherText[:aes.BlockSize]
	io.ReadFull(rand.Reader, iv)

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], source)

	return []byte(base64.StdEncoding.EncodeToString(cipherText))
}

func decrypt(source []byte) []byte {
	cipherText, _ := base64.StdEncoding.DecodeString(string(source))

	block, _ := aes.NewCipher(CRYPT_KEY)

	if len(cipherText) < aes.BlockSize {
		return nil
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return cipherText
}
