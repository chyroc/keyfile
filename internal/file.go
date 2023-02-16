package internal

import (
	"bytes"
	"fmt"
	"os"
)

func DecryptFile(path, account string) ([]byte, error) {
	secret, err := getSecret(account)
	if err != nil {
		return nil, err
	}
	bs, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file '%s' failed: %w", path, err)
	}
	bs = bytes.TrimSpace(bs)

	decryptData, err := AesDecrypt(bs, secret, true)
	if err != nil {
		return nil, fmt.Errorf("decrypt file '%s' failed: %w", path, err)
	}

	return decryptData, nil
}

func EncryptFile(path, account string) ([]byte, error) {
	secret, err := getSecret(account)
	if err != nil {
		return nil, err
	}
	bs, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file '%s' failed: %w", path, err)
	}
	bs = bytes.TrimSpace(bs)

	encryptData, err := AesEncrypt(bs, secret, true)
	if err != nil {
		return nil, fmt.Errorf("encrypt file '%s' failed: %w", path, err)
	}

	return encryptData, nil
}
