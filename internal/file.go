package internal

import "os"

func getOrInputSecret() {

}

func DecodeFile(path string) ([]byte, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	_ = bs
	return nil, err
}

func EncodeFile(path string) ([]byte, error) {
	return nil, nil
}
