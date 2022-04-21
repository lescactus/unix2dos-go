package cmd

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
)

// removeCRLF will remove the trailling carriage return and/or line feed in a byte slice if existing
func removeCRLF(data []byte) []byte {
	// If byte slice end with CRLF
	if len(data) > 0 && data[len(data)-2] == '\r' && data[len(data)-1] == '\n' {
		return data[0 : len(data)-2]
	}

	// If byte slice end with LF
	if len(data) > 0 && data[len(data)-1] == '\n' {
		return data[0 : len(data)-1]
	}
	return data
}

// convert will open and read file ain input and convert from Unix using Line Feed (LF)
// as newline into DOS/Windows text file by replacing the Line Feed with Carriage Return (CR)
// and Line Feed (LF) characters (CRLF)
func convert(file, direction string) error {
	// Open the input file
	in, err := os.Open(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer in.Close()

	// Open the output file
	out, err := os.OpenFile(output, os.O_RDWR|os.O_CREATE, fs.FileMode(mode))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer out.Close()

	// Create a buffered io reader
	reader := bufio.NewReader(in)
	for {
		// Read line by line until '\n' character
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Remove both CR+LF and/or single LF
		newLine := removeCRLF(line)

		switch direction {
		// If converting from LF to CRLF, add CRLF
		case "unix2dos":
			newLine = append(newLine, '\r', '\n')

		// If converting from CRLF to LF, add LF
		case "dos2unix":
			newLine = append(newLine, '\n')
		default:
			return fmt.Errorf("Error: unsupported direction. Must be 'unix2dos' or 'dos2unix'.")
		}

		_, err = out.Write(newLine)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	}

	return nil
}
