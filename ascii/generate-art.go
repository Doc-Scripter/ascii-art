package Ascii

import "strings"

// generateAsciiArt generates combined ASCII art from input string.
func GenerateAsciiArt(input string, asciiArtMap map[rune]string) string {
	// Split the input into runes
	runes := []rune(input)

	// Create a slice to hold the ASCII art lines for each rune
	var asciiArtLines [][]string

	// For each rune in the input, get its ASCII art and split it into lines
	for _, r := range runes {
		if art, ok := asciiArtMap[r]; ok {
			lines := strings.Split(art, "\n")
			asciiArtLines = append(asciiArtLines, lines)
		}
	}

	// Concatenate the ASCII art lines horizontally
	var result string
	for i := 0; i < len(asciiArtLines[0]); i++ {
		for _, lines := range asciiArtLines {
			if i < len(lines) {
				result += lines[i] + " "
			}
		}
		result += "\n"
	}

	return result
}
