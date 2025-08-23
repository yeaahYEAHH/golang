package cmd

import "github.com/spf13/cobra"

var unpackCMD = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file",
}

func init() {
	rootCMD.AddCommand(unpackCMD)
}
