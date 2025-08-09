package cmd

import "github.com/spf13/cobra"

var packCMD = &cobra.Command{
	Use:   "pack",
	Short: "Pack file",
}

func init() {
	rootCMD.AddCommand(packCMD)
}
