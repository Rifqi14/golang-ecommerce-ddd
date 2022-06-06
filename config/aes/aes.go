package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
)

type Credential struct {
	Key string
}

func (c *Credential) shaKey() (res []byte) {
	h := sha1.New()
	h.Write([]byte(c.Key))
	res = h.Sum(nil)
	res = res[:16]

	return res
}

// Encrypt
func (c *Credential) Encrypt(text string) (string, error) {
	key := c.shaKey()
	textByte := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	b := base64.StdEncoding.EncodeToString(textByte)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))

	return hex.EncodeToString(ciphertext), nil
}

// Decrypt
func (c *Credential) Decrypt(text string) (string, error) {
	key := c.shaKey()
	textByte, err := hex.DecodeString(text)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(textByte) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := textByte[:aes.BlockSize]
	textByte = textByte[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(textByte, textByte)

	data, err := base64.StdEncoding.DecodeString(string(textByte))
	if err != nil {
		return "", err
	}

	return string(data[:]), nil
}

// Decrypt string
func (c *Credential) DecryptString(text string) string {
	key := c.shaKey()
	ciphertext, err := hex.DecodeString(text)
	if err != nil {
		return ""
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}
	if len(ciphertext) < aes.BlockSize {
		return ""
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)
	data, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		return ""
	}

	return string(data[:])
}
