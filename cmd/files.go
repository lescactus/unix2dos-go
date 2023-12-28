package cmd

import (
	"fmt"
	"io/fs"
	"os"
)

func openFiles(file string) (*os.File, *os.File, error) {
	// Open the input file
	in, err := os.Open(file)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open input file %s for reading: %s", file, err)

	}

	// Open the output file
	out, err := os.OpenFile(output, os.O_RDWR|os.O_CREATE, fs.FileMode(mode))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open output file %s for writing: %s", output, err)
	}

	return in, out, nil
}
