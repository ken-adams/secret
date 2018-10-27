package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gophercises/secret/secret"
)

const filePath = "/Users/stefanlapcevic/go/bin/secrets"

func validateKey(key string) bool {
	return len(key) == 16
}

func emitError(message string) {
	fmt.Println("Error: ", message)
}

func main() {
	if len(os.Args) < 2 {
		emitError("Specify command you want to perform")
		return
	}
	flag.Parse()
	arguments := flag.Args()

	switch os.Args[1] {
	case "get":
		if len(arguments) != 3 {
			emitError("Invalid usage: get command")
			return
		}

		key := arguments[2]
		key = key[3:len(key)]

		if !validateKey(key) {
			emitError("Key is not valid- Length should be at least 16 characters")
			return
		}

		v, err := secret.FileVault(key, filePath)
		if err != nil {
			emitError(fmt.Sprint(err))
			return
		}
		value, err := v.Get(os.Args[2])
		if err != nil {
			emitError(fmt.Sprint(err))
			return
		}
		if len(*value) == 0 {
			emitError("No value in vault for given input")
		} else {
			fmt.Println(*value)
		}
		break
	case "set":
		if len(arguments) != 4 {
			emitError("Invalid usage: set command")
			return
		}

		key := arguments[3]
		key = key[3:len(key)]

		if !validateKey(key) {
			emitError("Key is not valid: Length should be at least 16 characters")
			return
		}

		v, err := secret.FileVault(key, filePath)
		if err != nil {
			emitError(fmt.Sprint(err))
			return
		}
		err = v.Set(arguments[1], arguments[2])
		if err != nil {
			emitError(fmt.Sprintf("Setting key failed: %s\n", err))
		}
		break
	default:
		emitError("Command not found!")
	}
}
