package cmd

import (
	"fmt"
	"log"

	"github.com/ken-adams/secret/secret/setcommand"
	"github.com/spf13/cobra"
)

var setKey *string
var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set value to vault for specific keyname",
	Long:  "Set value to vault for specific keyname",
}

func init() {
	setKey = SetCmd.Flags().StringP("keyname", "k", "", "Keyname you want to read")
}

func SetRunnerForSetCmd(runner setcommand.SetForKey) {
	SetCmd.Run = func(cmd *cobra.Command, args []string) {
		if err := runner(args[0], args[1], *setKey); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Value successfully set for key: %s\n", args[0])
	}
}
