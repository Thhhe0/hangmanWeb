package src

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func EncryptCBC(key, text []byte) (ciphertext []byte, err error) {
	for len(text)%aes.BlockSize != 0 {
		text = append(text, byte(' '))
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext = make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(ciphertext[aes.BlockSize:], text)

	return
}

func DecryptCBC(key, text []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	if len(text) < aes.BlockSize {
		fmt.Printf("ciphertext too short")
		return
	}

	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]

	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(text, text)

	plaintext = text

	return
}
