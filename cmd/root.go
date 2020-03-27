package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "",
	Short: "Secret binary application",
	Long:  "Secret binary application tool for reading cached file",
}

func init() {
	RootCmd.AddCommand(GetCmd)
	RootCmd.AddCommand(SetCmd)
}
