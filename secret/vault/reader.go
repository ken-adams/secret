package vault

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Vault map[string]string

type DecryptFn func(cipherstring string, keystring string) string
type EncryptFn func(plainstring, keystring string) (string, error)
type ReadVault func(key string) (*Vault, error)

type readFile func(filename string) ([]byte, error)

func VaultReader(filepath string, readEncrypted readEncrypt) ReadVault {
	return func(key string) (*Vault, error) {
		var result Vault

		decrypted, err := readEncrypted(key, filepath)
		if err != nil {
			return nil, errors.Wrapf(err, "Error while reading from file for path: %s", filepath)
		}

		if err = json.Unmarshal([]byte(decrypted), &result); err != nil {
			return nil, errors.Wrapf(err, "Error while converting file content")
		}

		return &result, nil
	}
}
