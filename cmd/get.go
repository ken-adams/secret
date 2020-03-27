package cmd

import (
	"fmt"
	"log"

	"github.com/ken-adams/secret/secret/getcommand"
	"github.com/spf13/cobra"
)

var getKey *string
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get value from vault",
	Long:  "Get value from vault for specific keyname",
}

func init() {
	getKey = GetCmd.Flags().StringP("keyname", "k", "", "Keyname you want to read")
}

func SetRunnerForGetCmd(runner getcommand.GetForKey) {
	GetCmd.Run = func(cmd *cobra.Command, args []string) {
		v, err := runner(args[0], *getKey)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Read value: %s\n", *v)
	}
}
