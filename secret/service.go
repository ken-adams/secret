package secret

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Secrets struct {
	Key        string
	PathToFile string
	Data       map[string]string
}

func (secrets *Secrets) Set(keyName, keyValue string) error {
	secrets.Data[keyName] = keyValue
	serialized, err := json.Marshal(secrets.Data)
	if err != nil {
		return errors.Wrapf(err, "Set- Cannot marhsal value")
	}
	encrypted, err := encrypt(string(serialized), secrets.Key)
	if err != nil {
		return errors.Wrapf(err, "Set- Encryption failed:")
	}
	err = writeToFile(encrypted, secrets.PathToFile)
	if err != nil {
		return errors.Wrapf(err, "Set- Cannot save value to vault")
	}
	return nil
}

func (secrets *Secrets) Get(keyName string) (*string, error) {
	value, ok := secrets.Data[keyName]
	if !ok {
		return nil, errors.New("Get- Cannot find key within vault")
	}
	return &value, nil
}

func FileVault(key, path_to_file string) (*Secrets, error) {
	encryptedContent, err := readFromFile(path_to_file)
	if err != nil {
		return nil, errors.Wrapf(err, "FileVault- Cannot read from file")
	}
	secrets := make(map[string]string)
	plaintext := "{}"
	if len(encryptedContent) > 0 {
		plaintext = decrypt(string(encryptedContent), key)
	}
	// fmt.Println(plaintext)
	err = json.Unmarshal([]byte(plaintext), &secrets)
	if err != nil {
		return nil, errors.Wrapf(err, "FileVault- Cannot unmarshal from flle")
	}
	return &Secrets{
		Key:        key,
		PathToFile: path_to_file,
		Data:       secrets,
	}, nil
}
