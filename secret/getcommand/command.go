package getcommand

import (
	"fmt"

	"github.com/ken-adams/secret/secret/vault"
)

type GetForKey func(string, string) (*string, error)
type setForKey func(string, string) error

func NewCommand(readVault vault.ReadVault) GetForKey {
	return func(keyName, key string) (*string, error) {
		vault, err := readVault(key)
		if err != nil {
			return nil, err
		}
		value, ok := (*vault)[keyName]
		if !ok {
			return nil, fmt.Errorf("get: cannot find value for key: %q", keyName)
		}
		return &value, nil
	}
}
