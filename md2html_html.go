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
)

// Shild HTML character
func shieldHTMLChar(char string) string {
	switch char {
	case "<":
		return "&lt"
	case ">":
		return "&gt"
	case "&":
		return "&amp"
	}

	return char
}

// Shields all HTML code inside text
func shieldHTML(text string) string {
	var result string = ""

	textRune := []rune(text)

	for i := range textRune {
		char := string(textRune[i])

		result = result + shieldHTMLChar(char)
	}

	return result
}

// Convert title string to HTML ID (<h1 id="......">)
func toHTMLID(line string) string {
	var result string = ""

	lineRune := []rune(line)

	var lastChar string = " "
	var tagOpen bool = true

	for i := range lineRune {
		char := string(lineRune[i])

		// HTML tags
		if char == "<" {
			tagOpen = true

		} else if char == ">" {
			tagOpen = false

		} else if tagOpen == true {
			//pass

			// Special chars
		} else if char == " " {
			if lastChar != " " {
				result = result + "_"
			}

		} else if char == "#" {
			//pass

		} else if char == "&" {
			//pass

		} else if char == "/" {
			//pass

		} else if char == `\` {
			//pass

		} else if char == "|" {
			//pass

		} else if char == ":" {
			//pass

		} else if char == ";" {
			//pass

		} else if char == `"` {
			//pass

		} else if char == "'" {
			//pass

		} else if char == "`" {
			//pass

		} else if char == "~" {
			//pass

			// Else: save char
		} else {
			result = result + char
		}

		lastChar = char
	}

	return strings.ToLower(result)
}
