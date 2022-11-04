package cli

import (
	"bufio"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const AppName = "stego"

var (
	Version = "0.0.0-dev"
	verbose bool
	silent  bool
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   AppName,
		Short: "stego",
		Long:  ``,
	}

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", false, "Silent mode (disable output)")

	rootCmd.AddCommand(
		newEncryptCmd(),
		newDecryptCmd(),
		newImagesCmd(),
		newVersionCmd(),
	)

	return rootCmd
}

func getInputFromStdin(cmd *cobra.Command) ([]byte, error) {
	fmt.Fprintf(cmd.OutOrStdout(), "Enter text: ")

	reader := bufio.NewReader(cmd.InOrStdin())

	text, err := reader.ReadBytes('\n')
	if err != nil {
		return nil, errors.Wrap(err, "failed reading bytes from stdin")
	}

	return text, nil
}
