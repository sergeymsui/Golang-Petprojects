package cmd

import (
	"errors"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Vlc programm",
	Run:   pack,
}

var (
	ErrEmptypath = errors.New("path to file is not specified")
)

func pack(_ *cobra.Command, args []string) {
	if len(args) < 2 || args[0] == "" || args[1] == "" {
		handleError(ErrEmptypath)
	}

	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleError(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		handleError(err)
	}

	packed := Encode(data)
	packedFileName := args[1]

	err = os.WriteFile(packedFileName, packed, 0644)

	if err != nil {
		handleError(err)
	}
}

func Encode(data []byte) []byte {

	s := "Hello World"
	arr := []byte{}

	for i := 0; i < len(s); i++ {
		arr = append(arr, s[i])
	}

	return arr
}

func init() {
	packCmd.AddCommand(vlcCmd)
}
