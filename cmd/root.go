package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use:   "arch",
	Short: "Simple archivator",
}

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		handleError(err)
	}
}

func handleError(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
