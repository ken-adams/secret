package setcommand

import (
	"encoding/json"
	"os"

	"github.com/ken-adams/secret/secret/vault"
	"github.com/pkg/errors"
)

type SetForKey func(string, string, string) error
type writeFileFn func(filename string, data []byte, perm os.FileMode) error

func NewCommand(filepath string, readVault vault.ReadVault, encrypt vault.EncryptFn, writeFile writeFileFn) SetForKey {
	return func(keyName, value, key string) error {
		vault, err := readVault(key)
		if err != nil {
			return err
		}

		(*vault)[keyName] = value
		ser, _ := json.Marshal(*vault)

		encrypted, err := encrypt(string(ser), key)
		if err != nil {
			return errors.Wrapf(err, "Set- Encryption failed:")
		}

		if err := writeFile(filepath, []byte(encrypted), 777); err != nil {
			return errors.Wrapf(err, "Set- Cannot save value to vault")
		}

		return nil
	}
}
