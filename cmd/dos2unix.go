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
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			in, out, err := openFiles(args[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			defer in.Close()
			defer out.Close()

			err = convert(in, out, dos2unix)
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
