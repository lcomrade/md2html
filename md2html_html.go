// Copyright (C) 2022 Leonid Maslakov.

// This file is part of md2html.

// md2html is free software: you can redistribute it
// and/or modify it under the terms of the
// GNU Affero Public License as published by the
// Free Software Foundation, either version 3 of the License,
// or (at your option) any later version.

// md2html is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
// or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero Public License for more details.

// You should have received a copy of the GNU Affero Public License along with md2html.
// If not, see <https://www.gnu.org/licenses/>.

package md2html

import (
	"strings"
	"unicode"
)

// Shild HTML character
func shieldHTMLChar(char rune) string {
	switch char {
	case '<':
		return "&lt"
	case '>':
		return "&gt"
	case '&':
		return "&amp"
	}

	return string(char)
}

// Shields all HTML code inside text
func shieldHTML(text string) string {
	var result string = ""

	textRune := []rune(text)

	for i := range textRune {
		charRune := textRune[i]

		result = result + shieldHTMLChar(charRune)
	}

	return result
}

// Convert title string to HTML ID (<h1 id="......">)
func toHTMLID(line string) string {
	var result string = ""

	lineRune := []rune(line)

	var tagOpen bool = false
	var lastLetter bool = false

	for i := range lineRune {
		charRune := lineRune[i]
		char := string(charRune)

		// HTML tags
		if char == "<" {
			tagOpen = true

		} else if char == ">" {
			tagOpen = false

		} else if tagOpen == true {
			//pass

			// Special chars
		} else if unicode.IsLetter(charRune) == false {
			if lastLetter == true {
				lastLetter = false
				result = result + "-"
			}

			// Else: save char
		} else {
			lastLetter = true
			result = result + char
		}
	}

	// Remove '-' from beginning and end
	if result[0] == '-' {
		result = string(result[1:])
	}

	resultLen := len(result)
	if result[resultLen-1] == '-' {
		result = string(result[:resultLen-1])
	}

	return strings.ToLower(result)
}
