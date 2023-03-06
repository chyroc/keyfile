package internal

import (
	"bytes"
	"fmt"
	"os"
	"time"
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

func WaitFileChanged(filePath string) (chan bool, error) {
	initialStat, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	res := make(chan bool, 1)
	var finalErr error

	go func() {
		for {
			if finalErr != nil {
				return
			}
			stat, err := os.Stat(filePath)
			if err != nil {
				finalErr = err
				return
			}
			if stat.Size() != initialStat.Size() || stat.ModTime() != initialStat.ModTime() {
				res <- true
				return
			}
			time.Sleep(time.Second / 2)
		}
	}()

	if finalErr != nil {
		res <- false
	}

	return res, finalErr
}

func WriteTempFile(bs []byte) (string, error) {
	tmpFile, err := os.CreateTemp("", "keyfile-*")
	if err != nil {
		return "", err
	}
	if _, err = tmpFile.Write(bs); err != nil {
		return "", err
	}
	_ = tmpFile.Close()
	return tmpFile.Name(), nil
}
