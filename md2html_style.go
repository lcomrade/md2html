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
	"unicode"
)

// Replace * to <em>
// Replace ** to <strong>
// Replace *** to <strong> + <em>
// Replace ~~ to <del>
// Replace ` to <code>
// Replace \* to *
func mdStyle(line string) string {
	var result string = ""

	lineRune := []rune(line)
	lineLen := len(lineRune)

	// Track opened HTML tags
	var emTagOpen bool = false
	var strongTagOpen bool = false
	var strongEmTagOpen bool = false
	var delTagOpen bool = false
	var codeTagOpen bool = false

	var skip int = 0

	// Reading line by character
	for i := range lineRune {
		// Skip
		if skip != 0 {
			skip = skip - 1
			continue
		}

		lastCharRune := ' '
		lastChar := " "
		charRune := lineRune[i]
		char := string(charRune)
		nextCharRune := ' '
		nextChar := " "
		nextNextCharRune := ' '
		nextNextChar := " "
		nextNextNextCharRune := ' '
		nextNextNextChar := " "

		// Get last char
		if i != 0 {
			lastCharRune = lineRune[i-1]
			lastChar = string(lastCharRune)
		}

		// Get next char
		if lineLen > i+1 {
			nextCharRune = lineRune[i+1]
			nextChar = string(nextCharRune)
		}

		// Get next next char
		if lineLen > i+2 {
			nextNextCharRune = lineRune[i+2]
			nextNextChar = string(nextNextCharRune)
		}

		// Get next next char
		if lineLen > i+3 {
			nextNextNextCharRune = lineRune[i+3]
			nextNextNextChar = string(nextNextNextCharRune)
		}

		// Shielding characters inside <code>....</code>
		if char != "`" && codeTagOpen == true {
			result = result + shieldHTMLChar(charRune)

			// Remove many spaces
		} else if lastChar == " " && char == " " {
			//pass

			// Replace \` to `
		} else if char == `\` && nextChar == "`" {
			result = result + "`"
			skip = 1

			// Replace \* to *
		} else if char == `\` && nextChar == "*" {
			result = result + "*"
			skip = 1

			// Replace \_ to _
		} else if char == `\` && nextChar == "_" {
			result = result + "_"
			skip = 1

			// Replace \~ to ~
		} else if char == `\` && nextChar == "~" {
			result = result + "~"
			skip = 1

			// Replace \# to #
		} else if char == `\` && nextChar == "#" {
			result = result + "#"
			skip = 1

			// Replace \+ to +
		} else if char == `\` && nextChar == "+" {
			result = result + "+"
			skip = 1

			// Replace \- to -
		} else if char == `\` && nextChar == "-" {
			result = result + "-"
			skip = 1

			// Replace \. to .
		} else if char == `\` && nextChar == "." {
			result = result + "."
			skip = 1

			// ^ - space
			// Bold and italic text
		} else if char == "*" || char == "_" {
			// ^*^
			if lastChar == " " && nextChar == " " {
				result = result + char

				// a*a
			} else if unicode.IsLetter(lastCharRune) && unicode.IsLetter(nextCharRune) {
				result = result + char

				// ^***WORD.... or ....WORD***^
			} else if lastChar != char && nextChar == char && nextNextChar == char && nextNextNextChar != char {
				//} else if unicode.IsLetter(lastCharRune) == false && nextChar == char && nextNextChar == char && nextNextNextChar != char {
				if strongEmTagOpen == false {
					if strongTagOpen == false {
						result = result + "<strong>"
						strongTagOpen = true
					}

					if emTagOpen == false {
						result = result + "<em>"
						emTagOpen = true
					}

					strongEmTagOpen = true

				} else {
					if emTagOpen == true {
						result = result + "</em>"
						emTagOpen = false
					}

					if strongTagOpen == true {
						result = result + "</strong>"
						strongTagOpen = false
					}

					strongEmTagOpen = false
				}

				skip = 2

				// ^**WORD.... or ....WORD**^
			} else if lastChar != char && nextChar == char && nextNextChar != char {
				if strongTagOpen == false {
					result = result + "<strong>"
					strongTagOpen = true

				} else {
					result = result + "</strong>"
					strongTagOpen = false
				}

				skip = 1

				// ^*WORD.... or ....WORD*^
			} else if lastChar != char && nextChar != char {
				if emTagOpen == false {
					result = result + "<em>"
					emTagOpen = true

				} else {
					result = result + "</em>"
					emTagOpen = false
				}
			}

			// ^ - space
			// Strikethrough text
		} else if char == "~" {
			// ^~^
			if lastChar == " " && nextChar == " " {
				result = result + "~"

				// a~a
			} else if unicode.IsLetter(lastCharRune) && unicode.IsLetter(nextCharRune) {
				result = result + "~"

				// ^~~WORD.... or ....WORD~~^
			} else if nextChar == "~" {
				if delTagOpen == false {
					result = result + "<del>"
					delTagOpen = true

				} else {
					result = result + "</del>"
					delTagOpen = false
				}

				skip = 1

			} else {
				result = result + "~"
			}

			// ^ - space
			// Code quote
		} else if char == "`" {
			// a`a
			if unicode.IsLetter(lastCharRune) && unicode.IsLetter(nextCharRune) {
				result = result + "`"

				// ^`.... or ....`^
			} else {
				if codeTagOpen == false {
					result = result + "<code>"
					codeTagOpen = true

				} else {
					result = result + "</code>"
					codeTagOpen = false
				}
			}

			// If not formated text
		} else {
			result = result + char
		}
	}

	// Remove space from end
	resultLen := len(result)
	if result[resultLen-1] == ' ' {
		result = string(result[:resultLen-1])
	}

	// If HTML tags not closed
	if emTagOpen == true {
		result = result + "</em>"
		emTagOpen = false
	}

	if strongTagOpen == true {
		result = result + "</strong>"
		strongTagOpen = false
	}

	if delTagOpen == true {
		result = result + "</del>"
		delTagOpen = false
	}

	if codeTagOpen == true {
		result = result + "</code>"
		codeTagOpen = false
	}

	// Return result
	return result
}
