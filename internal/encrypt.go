package internal

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func pkcs5Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func paddingKey(key []byte) []byte {
	keyLen := len(key)
	if keyLen < 16 {
		key = append(key, bytes.Repeat([]byte{0}, 16-keyLen)...)
	} else if keyLen < 24 {
		key = append(key, bytes.Repeat([]byte{0}, 24-keyLen)...)
	} else if keyLen < 32 {
		key = append(key, bytes.Repeat([]byte{0}, 32-keyLen)...)
	}
	return key
}

func AesEncrypt(originData, key []byte, withBase64 bool) ([]byte, error) {
	key = paddingKey(key)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	originData = pkcs5Padding(originData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	encryptData := make([]byte, len(originData))
	blockMode.CryptBlocks(encryptData, originData)

	if withBase64 {
		return []byte(base64.StdEncoding.EncodeToString(encryptData)), nil
	}

	return encryptData, nil
}

func AesDecrypt(encryptData, key []byte, withBase64 bool) ([]byte, error) {
	if withBase64 {
		var err error
		encryptData, err = base64.StdEncoding.DecodeString(string(encryptData))
		if err != nil {
			return nil, err
		}
	}
	key = paddingKey(key)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	originData := make([]byte, len(encryptData))
	blockMode.CryptBlocks(originData, encryptData)
	originData = pkcs5UnPadding(originData)
	return originData, nil
}
