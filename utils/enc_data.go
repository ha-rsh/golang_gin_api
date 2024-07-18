package utils

import (
	"io"
	"errors"
	"crypto/aes"
	"crypto/rand"
	"crypto/cipher"
	"encoding/base64"
)


func Decrypt(key []byte, encrypted string) (string, error) {
    ciphertext, err := base64.RawURLEncoding.DecodeString(encrypted)
    if err != nil {
        return "nil", err
    }
    block, err := aes.NewCipher(key)
    if err != nil {
        return "nil", err
    }
    if len(ciphertext) < aes.BlockSize {
        return "nil", errors.New("ciphertext too short")
    }
    iv := ciphertext[:aes.BlockSize]
    ciphertext = ciphertext[aes.BlockSize:]
    cfb := cipher.NewCFBDecrypter(block, iv)
    cfb.XORKeyStream(ciphertext, ciphertext)
    return string(ciphertext), nil
}

func Encrypt(key, data []byte) (string, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }
    ciphertext := make([]byte, aes.BlockSize+len(data))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }
    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], data)
    return base64.RawURLEncoding.EncodeToString(ciphertext), nil
}

