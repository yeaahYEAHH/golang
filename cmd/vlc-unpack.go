package cmd

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"archivator/lib/vlc"

	"github.com/spf13/cobra"
)

var unpackedExtention string = "txt"

var vlcUnpackCMD = &cobra.Command{
	Use:   "vlc",
	Short: "method packing file by Variable-length code",
	Run:   unpack,
}

func unpack(_ *cobra.Command, args []string) {
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

	packed := vlc.Decode(string(data))

	if err := os.WriteFile(unpackedFilename(path), []byte(packed), 0644); err != nil {
		handleError(err)
	}
}

func unpackedFilename(path string) string {
	filename := filepath.Base(path)
	ext := filepath.Ext(filename)

	return strings.TrimSuffix(filename, ext) + "." + unpackedExtention
}

func init() {
	unpackCMD.AddCommand(vlcUnpackCMD)
}
