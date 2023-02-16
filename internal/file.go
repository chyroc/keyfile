package internal

import (
	"fmt"
	"os"
)

func DecodeFile(path, account string) ([]byte, error) {
	secret, err := getSecret(account)
	if err != nil {
		return nil, err
	}
	bs, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file '%s' failed: %w", path, err)
	}

	decryptData, err := AesDecrypt(bs, secret)
	if err != nil {
		return nil, fmt.Errorf("decrypt file '%s' failed: %w", path, err)
	}

	return decryptData, nil
}

func EncodeFile(path, account string) ([]byte, error) {
	secret, err := getSecret(account)
	if err != nil {
		return nil, err
	}
	bs, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file '%s' failed: %w", path, err)
	}
	encryptData, err := AesEncrypt(bs, secret)
	if err != nil {
		return nil, fmt.Errorf("encrypt file '%s' failed: %w", path, err)
	}

	return encryptData, nil
}
