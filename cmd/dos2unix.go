package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	dos2unixCmd = &cobra.Command{
		Use:   "dos2unix",
		Short: "Convert DOS file format to Unix file format",
		Long: `dos2unix will convert DOS/Windows files using Carriage Return and Line Feed (CRLF) as newline into Unix text file by replacing the Carriage Return Line Feed with
Line Feed (LF) character.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				fmt.Fprintf(os.Stderr, "Error: expected 1 argument, got %d\n", len(args))
				os.Exit(1)
			}

			err := convert(args[0], "dos2unix")
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(dos2unixCmd)
}
