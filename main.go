package main

import (
	"flag"
	"fmt"

	"github.com/gophercises/secret/secret"
)

const filePath = "/Users/stefanlapcevic/go/bin/secrets"

func validateKey(key string) bool {
	return len(key) == 16
}

func parseKey(arg string) *string {
	var result *string
	if len(arg) > 2 && arg[0:3] == "-k=" {
		value := arg[3:len(arg)]
		result = &value
	}
	return result
}

func emitError(message string) {
	fmt.Println("Error: ", message)
}

func main() {
	flag.Parse()
	arguments := flag.Args()

	if len(arguments) < 1 {
		emitError("Specify command you want to perform")
		return
	}

	switch arguments[0] {
	case "get":
		if len(arguments) != 3 {
			emitError("Invalid usage: get command")
			return
		}

		key := parseKey(arguments[2])
		if key == nil {
			emitError("Invalid usage: get command")
			return
		}

		if !validateKey(*key) {
			emitError("Key is not valid- Length should be at least 16 characters")
			return
		}

		v, err := secret.FileVault(*key, filePath)
		if err != nil {
			emitError(fmt.Sprint(err))
			return
		}
		value, err := v.Get(arguments[1])
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

		key := parseKey(arguments[3])
		if key == nil {
			emitError("Invalid usage: set command")
			return
		}

		if !validateKey(*key) {
			emitError("Key is not valid: Length should be at least 16 characters")
			return
		}

		v, err := secret.FileVault(*key, filePath)
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
