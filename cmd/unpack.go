package cmd

import (
	"archivator/lib/compression"
	"archivator/lib/compression/vlc"
	"archivator/lib/compression/vlc/table/shannon_fano"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var unpackedExtention string = "txt"

var unpackCMD = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file",
	Run:   unpack,
}

func unpack(cmd *cobra.Command, args []string) {
	var decoder compression.Decoder

	if len(args) == 0 || args[0] == "" {
		handleError(emptyError)
	}

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		decoder = vlc.NewEncoderDecoder(shannon_fano.Generator{})
	default:
		cmd.PrintErrf("Unsupported method: %s", method)
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

	packed := decoder.Decode(data)

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
	rootCMD.AddCommand(unpackCMD)

	unpackCMD.Flags().StringP("method", "m", "", "decompression method to use: vlc")

	if err := unpackCMD.MarkFlagRequired("method"); err != nil {
		handleError(err)
	}
}
