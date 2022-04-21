package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	unix2dosCmd = &cobra.Command{
		Use:   "unix2dos",
		Short: "Convert Unix file format to DOS file format",
		Long: `unix2dos will convert Unix files using Line Feed (LF) as newline into DOS/Windows text file by replacing the Line Feed with
Carriage Return (CR) and Line Feed (LF) characters (CRLF).`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				fmt.Fprintf(os.Stderr, "Error: expected 1 argument, got %d\n", len(args))
				os.Exit(1)
			}

			err := convert(args[0], "unix2dos")
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(unix2dosCmd)
}
