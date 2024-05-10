package main

import (
	"fmt"
	"os"
	"strings"

	"art/Ascii"
)

func main() {
	// Check command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <input string>")
		os.Exit(1)
	}

	// Read the ASCII art from the file
	asciiArtMap, err := Ascii.ReadAscii("standard.txt")
	if len(asciiArtMap)-1 != 126-32 {
		fmt.Println("ascii characters are less")
		os.Exit(1)
	}
	if err != nil {
		fmt.Println("Error reading ASCII art file:", err)
		os.Exit(1)
	}

	// Process each argument
	for _, arg := range os.Args[1:] {
		if arg == "" {
			fmt.Println("Usage: go run . <input string>")
			os.Exit(1)
		}
		result := strings.Split(arg, "\\n")
		for _, art := range result {
			art := Ascii.GenerateAsciiArt(art, asciiArtMap)
			fmt.Print(art)
		}
	}
}
