package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	mode   uint32
	output string

	rootCmd = &cobra.Command{
		Use:   "unix2dos-go",
		Short: "DOS & Mac to Unix and vice versa text file format converter",
		Long: `unix2dos-go is a simple cli written in go to convert line breaks in a text file from Unix format (LF) to DOS format (CR+LF) and vice versa.
In DOS/Windows text files a line break, also known as newline, is a combination of two characters: a Carriage Return (CR) followed by a Line Feed (LF). In Unix text files a line break is a single character: the Line Feed
(LF). In Mac text files, prior to Mac OS X, a line break was single Carriage Return (CR) character. Nowadays Mac OS uses Unix style (LF) line breaks.

It is inspired by the unix2dos utility.


The source code is available at https://github.com/lescactus/unix2dos-go
`,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Uint32VarP(&mode, "mode", "m", 0644, "Unix permission numeric mode in octal for the converted file. See chmod(1) for more informations.")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "./unix2dos.converted", "Name of the converted file.")
}
