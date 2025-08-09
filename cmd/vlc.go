package cmd

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var packedExtention string = "vlc"

var emptyError = errors.New("path to filt is not sprcified")

var vlcCMD = &cobra.Command{
	Use:   "vlc",
	Short: "method packing file by Variable-length code",
	Run:   vlc,
}

func vlc(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleError(emptyError)
	}

	path := args[0]

	read, err := os.Open(path)
	if err != nil {
		handleError(err)
	}
	defer read.Close()

	data, err := io.ReadAll(read)
	if err != nil {
		handleError(err)
	}

	packed := data

	if err := os.WriteFile(packedFilename(path), packed, 0644); err != nil {
		handleError(err)
	}
}

func packedFilename(path string) string {
	filename := filepath.Base(path)
	ext := filepath.Ext(filename)

	return strings.TrimSuffix(filename, ext) + "." + packedExtention
}

func init() {
	packCMD.AddCommand(vlcCMD)
}
