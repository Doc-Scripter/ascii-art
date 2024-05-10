package Ascii

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

// readAscii reads ASCII art from the file and returns a map of runes to art strings.
func ReadAscii(filename string) (map[rune]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	asciiArtMap := make(map[rune]string)
	scanner := bufio.NewScanner(file)

	var currentRune rune = 32 // Start from space character
	var artLines []string
	var lineCount int
	var maxLen int

	for scanner.Scan() {
		line := scanner.Text()
		artLines = append(artLines, line)
		lineCount++

		// Find the length of the line up to the last character plus one space
		lastCharIndex := strings.LastIndexFunc(line, func(c rune) bool {
			return !unicode.IsSpace(c)
		})
		if lastCharIndex+1 > maxLen {
			maxLen = lastCharIndex + 1
		}

		// Every 9 lines, map the ASCII art to the current rune and move to the next rune
		if lineCount%9 == 0 {
			for _, artLine := range artLines[1:] {
				asciiArtMap[currentRune] += artLine[:maxLen] + "\n"
			}
			artLines = artLines[:0] // Reset artLines
			currentRune++           // Move to the next rune
			maxLen = 0              // Reset maxLen
			if currentRune > 126 {  // Stop if we've reached the tilde (~)
				break
			}
		}
	}

	return asciiArtMap, scanner.Err()
}
