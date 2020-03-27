package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/ken-adams/secret/cmd"
	"github.com/ken-adams/secret/secret"
	"github.com/ken-adams/secret/secret/getcommand"
	"github.com/ken-adams/secret/secret/setcommand"
	"github.com/ken-adams/secret/secret/vault"
)

const filePath = "./secret_"

func main() {
	// create dummy loger for testing
	logger := log.New(os.Stdout, "Secret Logger: ", log.Lshortfile)

	// initialize vault reader
	readVault := vault.VaultReader(
		filePath,
		vault.NewReadEncrypt(
			ioutil.ReadFile,
			secret.Decrypt,
		).Log(logger),
	)

	cmd.SetRunnerForGetCmd(getcommand.NewCommand(readVault))
	cmd.SetRunnerForSetCmd(setcommand.NewCommand(filePath, readVault, secret.Encrypt, ioutil.WriteFile))
	cmd.RootCmd.Execute()
}
