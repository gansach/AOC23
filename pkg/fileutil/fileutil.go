// pkg/fileutil/fileutil.go
package fileutil

import (
  "fmt"
	"bufio"
	"os"
)

func ReadFileLines(filename string) ([]string, error) {
	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close() // Ensure the file is closed when done

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var lines []string

	// Iterate through each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return lines, nil
}
