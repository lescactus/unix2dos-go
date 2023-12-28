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
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			in, out, err := openFiles(args[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			defer in.Close()
			defer out.Close()

			err = convert(in, out, unix2dos)
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
