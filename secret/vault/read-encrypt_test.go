package vault

import (
	"errors"
	"testing"

	"gotest.tools/assert"
)

func TestNewReadEncrypt(t *testing.T) {
	cases := []struct {
		read     readFile
		decrypt  DecryptFn
		key      string
		filepath string
		output   string
		err      error
	}{
		{
			read:     func(filename string) ([]byte, error) { return []byte(""), nil },
			decrypt:  func(cipherstring string, keystring string) string { return cipherstring },
			key:      "KEY-1",
			filepath: "testfile",
			output:   "{}",
		},
		{
			read:     func(filename string) ([]byte, error) { return []byte("{'test':'test'}"), nil },
			decrypt:  func(cipherstring string, keystring string) string { return cipherstring },
			key:      "KEY-1",
			filepath: "testfile",
			output:   "{'test':'test'}",
		},
		{
			read:     func(filename string) ([]byte, error) { return []byte(""), errors.New("read failed") },
			decrypt:  func(cipherstring string, keystring string) string { return cipherstring },
			key:      "KEY-1",
			filepath: "testfile",
			output:   "{'test':'test'}",
			err:      errors.New("Error while reading from file for path: testfile: read failed"),
		},
	}
	for idx, c := range cases {
		res, err := NewReadEncrypt(c.read, c.decrypt)(c.key, c.filepath)
		if c.err != nil {
			assert.Error(t, err, c.err.Error(), "case: %d", idx)
			continue
		}
		assert.Equal(t, res, c.output, "case: %d", idx)
	}
}
