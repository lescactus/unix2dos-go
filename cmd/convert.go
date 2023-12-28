package cmd

import (
	"bufio"
	"fmt"
	"io"
)

type Direction string

const (
	unix2dos = Direction("unix2dos")
	dos2unix = Direction("dos2unix")
)

// removeCRLF will remove the trailling carriage return and/or line feed in a byte slice if existing
func removeCRLF(data []byte) []byte {
	// If byte slice ends with CRLF
	if len(data) > 1 && data[len(data)-2] == '\r' && data[len(data)-1] == '\n' {
		return data[0 : len(data)-2]
	}

	// If byte slice ends with LF
	if len(data) > 0 && data[len(data)-1] == '\n' {
		return data[0 : len(data)-1]
	}
	return data
}

// convert will read from the given io.Reader and convert from Unix using Line Feed (LF)
// as newline into DOS/Windows text by replacing the Line Feed with Carriage Return (CR)
// and Line Feed (LF) characters (CRLF) and vice-versa
func convert(in io.Reader, out io.Writer, direction Direction) error {
	reader := bufio.NewReader(in)

	for {
		// Read line by line until '\n' character
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read bytes: %w", err)
		}

		// Remove both CR+LF and/or single LF
		newLine := removeCRLF(line)

		switch direction {
		// If converting from LF to CRLF, add CRLF
		case unix2dos:
			newLine = append(newLine, '\r', '\n')

		// If converting from CRLF to LF, add LF
		case dos2unix:
			newLine = append(newLine, '\n')
		default:
			return fmt.Errorf("error: unsupported direction. Must be 'unix2dos' or 'dos2unix'")
		}

		_, err = out.Write(newLine)
		if err != nil {
			return fmt.Errorf("failed to write to buffer: %w", err)
		}
	}

	return nil
}
