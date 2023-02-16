package internal

import (
	"github.com/keybase/go-keychain"
	"github.com/pkg/errors"
)

var ErrDuplicateItem = keychain.ErrorDuplicateItem
var ErrItemNotFound = keychain.ErrorItemNotFound

func AddKeyChain(account string, data []byte) error {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(service)
	item.SetAccount(account)
	item.SetLabel(account)
	item.SetData(data)
	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetAccessible(keychain.AccessibleAfterFirstUnlockThisDeviceOnly)
	err := keychain.AddItem(item)
	if errors.Is(err, keychain.ErrorDuplicateItem) {
		return ErrDuplicateItem
	}

	return err
}

func GetKeyChain(account string) ([]byte, error) {
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassGenericPassword)
	query.SetService(service)
	query.SetAccount(account)
	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnData(true)
	results, err := keychain.QueryItem(query)
	if err != nil {
		return nil, err
	} else if len(results) == 0 {
		return nil, ErrItemNotFound
	}
	return results[0].Data, nil
}
