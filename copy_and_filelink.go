package main

// this is for a linux system
//when a user right-clicks a file and selects "Copy as Path" this code handles
// copying file-path and creating a file:// link

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
)

// Writes a file:// link (URL-encoded) to dst
// i would like to have it as a hyperlink also
func create_file_hyperlink(filePath, linkPath string) error {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}
	u := url.URL{
		Scheme: "file",
		Path:   absPath,
	}
	// Use EscapedPath to ensure umlauts and unicode are encoded, but slashes are preserved
	fileURL := "file://" + u.EscapedPath()
	linkText := fmt.Sprintf("[%s](%s)", filepath.Base(absPath), fileURL)
	return os.WriteFile(linkPath, []byte(linkText), 0644)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input_file> <output_file>\n", os.Args[0])
		os.Exit(1)
	}
	input := os.Args[1]
	output := os.Args[2]

	err := create_file_hyperlink(input, output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
