package vault

import (
	"errors"
	"testing"

	"gotest.tools/assert"
)

func TestVaultReader(t *testing.T) {
	cases := []struct {
		readEncrypted readEncrypt
		key           string
		output        Vault
		err           error
	}{
		{
			readEncrypted: func(key, filepath string) (string, error) {
				return `{"first":"first","second":"second"}`, nil
			},
			key:    "key",
			output: Vault{"first": "first", "second": "second"},
		},
		{
			readEncrypted: func(key, filepath string) (string, error) { return `{}`, nil },
			key:           "key",
			output:        Vault{},
		},
		{
			readEncrypted: func(key, filepath string) (string, error) { return `{}`, errors.New("dummy") },
			key:           "key",
			err:           errors.New("Error while reading from file for path: filename: dummy"),
		},
	}
	for idx, c := range cases {
		res, err := VaultReader("filename", c.readEncrypted)(c.key)
		if err != nil {
			assert.Error(t, err, c.err.Error(), "case: %d", idx)
			continue
		}
		if res == nil {
			t.Fatalf("vault expected not to be nil for case: %d", idx)
		}
		assert.DeepEqual(t, *res, c.output)
	}
}
