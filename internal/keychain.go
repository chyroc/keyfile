package internal

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/keybase/go-keychain"
)

var (
	ErrDuplicateItem = keychain.ErrorDuplicateItem
	ErrItemNotFound  = keychain.ErrorItemNotFound
)

func AddKeyChain(account string, data []byte) error {
	item := keychain.NewGenericPassword(service, account, account, data, "")
	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetAccessible(keychain.AccessibleAfterFirstUnlockThisDeviceOnly)
	err := keychain.AddItem(item)
	if errors.Is(err, keychain.ErrorDuplicateItem) {
		return ErrDuplicateItem
	}

	return err
}

func GetKeyChain(account string) ([]byte, error) {
	item := keychain.NewGenericPassword(service, account, account, nil, "")
	item.SetMatchLimit(keychain.MatchLimitOne)
	item.SetReturnData(true)
	results, err := keychain.QueryItem(item)
	if err != nil {
		return nil, err
	} else if len(results) == 0 {
		return nil, ErrItemNotFound
	}
	return results[0].Data, nil
}

func DeleteKeyChain(account string) error {
	item := keychain.NewGenericPassword(service, account, account, nil, "")
	return keychain.DeleteItem(item)
}

func SetKeyChain(account string, data []byte) error {
	if len(data) == 0 {
		return fmt.Errorf("set keychain error: data is empty")
	}

	bs, err := GetKeyChain(account)
	if err != nil && !errors.Is(err, ErrItemNotFound) {
		return err
	} else if errors.Is(err, ErrItemNotFound) {
		return AddKeyChain(account, data)
	} else if bytes.Equal(bs, data) {
		return nil
	}

	where := keychain.NewGenericPassword(service, account, account, nil, "")
	updateItem := keychain.NewGenericPassword(service, account, account, data, "")

	return keychain.UpdateItem(where, updateItem)
}
