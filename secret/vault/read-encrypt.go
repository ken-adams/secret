package vault

import (
	"log"

	"github.com/pkg/errors"
)

type readEncrypt func(key, filepath string) (string, error)

func (f readEncrypt) Log(logger *log.Logger) readEncrypt {
	return func(key, filepath string) (string, error) {
		res, err := f(key, filepath)
		if err != nil {
			logger.Printf("error reading/encrypting: %s", err)
			return "", err
		}
		logger.Printf("value successfully read and decrypted. File: %s\n", res)
		return res, err
	}
}

func NewReadEncrypt(read readFile, decrypt DecryptFn) readEncrypt {
	return func(key, filepath string) (string, error) {
		plaintext := "{}"
		encrypted, err := read(filepath)
		if err != nil {
			return "", errors.Wrapf(err, "Error while reading from file for path: %s", filepath)
		}

		if len(encrypted) > 0 {
			plaintext = decrypt(string(encrypted), key)
		}
		return plaintext, nil
	}
}
