package internal

import (
	"errors"
	"fmt"
)

func getSecret(account string) ([]byte, error) {
	account, err := inputText("account", account)
	if err != nil {
		return nil, err
	}

	data, err := GetKeyChain(account)
	if err != nil && !errors.Is(err, ErrItemNotFound) {
		return nil, err
	} else if len(data) != 0 {
		if len(data) > 32 {
			return nil, fmt.Errorf("large secret size: %d", len(data))
		}
		return data, nil
	}

	secret, err := inputText("secret", "")
	if err != nil {
		return nil, err
	}

	if err = AddKeyChain(account, []byte(secret)); err != nil && !errors.Is(err, ErrDuplicateItem) {
		return nil, err
	}

	data, err = GetKeyChain(account)
	if err != nil {
		return nil, fmt.Errorf("get '%s' secret error: %w", account, err)
	} else if len(data) == 0 {
		return nil, fmt.Errorf("get '%s' secret error: save failed", account)
	} else if len(data) > 32 {
		return nil, fmt.Errorf("large secret size: %d", len(data))
	}

	return data, nil
}

func inputText(msg string, defaultValue string) (string, error) {
	if defaultValue != "" {
		return defaultValue, nil
	}
	fmt.Println("Please input", msg)
	var s string
	_, _ = fmt.Scanln(&s)
	if s == "" {
		return "", fmt.Errorf("get '%s' input empty", msg)
	}
	return s, nil
}
