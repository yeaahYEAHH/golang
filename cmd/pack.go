package cmd

import (
	"archivator/lib/compression"
	"archivator/lib/compression/vlc"
	"archivator/lib/compression/vlc/table/shannon_fano"
	"errors"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var packCMD = &cobra.Command{
	Use:   "pack",
	Short: "Pack file",
	Run:   pack,
}

var packedExtention string = "arch"

var emptyError = errors.New("path to file is not exist")

func pack(cmd *cobra.Command, args []string) {
	var encoder compression.Encoder

	if len(args) == 0 || args[0] == "" {
		handleError(emptyError)
	}

	method := cmd.Flag("method").Value.String()

	switch method {
	case "sf":
		encoder = vlc.NewEncoderDecoder(shannon_fano.Generator{})
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

	packed := encoder.Encode(string(data))

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
	rootCMD.AddCommand(packCMD)

	packCMD.Flags().StringP("method", "m", "", "compression method to use: sf")

	if err := packCMD.MarkFlagRequired("method"); err != nil {
		handleError(err)
	}
}
