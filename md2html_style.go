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
	var delTagOpen bool = false
	var codeTagOpen bool = false

	// Reading line by character
	for i := range lineRune {
		lastCharRune := ' '
		lastChar := " "
		char := string(lineRune[i])
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
			nextCharRune = lineRune[i+2]
			nextNextChar = string(nextCharRune)
		}

		// Get next next char
		if lineLen > i+3 {
			nextNextNextCharRune = lineRune[i+3]
			nextNextNextChar = string(nextNextNextCharRune)
		}

		// Shielding characters inside <code>....</code>
		if char != "`" && codeTagOpen == true {
			result = result + shieldHTMLChar(char)

			// Replace \* to *
		} else if char == `\` && nextChar == "*" {
			//pass

			// Replace \* to *
		} else if lastChar == `\` && char == "*" {
			result = result + "*"

			// Replace \_ to _
		} else if char == `\` && nextChar == "_" {
			//pass

			// Replace \_ to _
		} else if lastChar == `\` && char == "_" {
			result = result + "_"

			// Replace \~ to ~
		} else if lastChar == `\` && char == "~" {
			result = result + "~"

			// Replace \~ to ~
		} else if char == `\` && nextChar == "~" {
			//pass

			// Replace \~ to ~
		} else if lastChar == `\` && char == "~" {
			result = result + "~"

			// Replace \. to .
		} else if char == `\` && nextChar == "." {
			//pass

			// Replace \. to .
		} else if lastChar == `\` && char == "." {
			result = result + "."

			// ^ - space
			// Bold and italic text
		} else if char == "*" || char == "_" {
			// ^*^
			if lastChar == " " && nextChar == " " {
				result = result + char

				// a*a
			} else if unicode.IsLetter(lastCharRune) && unicode.IsLetter(nextCharRune) {
				result = result + char

				// ^***WORD....
			} else if unicode.IsLetter(lastCharRune) == false && nextChar == char && nextNextChar == char && nextNextNextChar != char {
				if strongTagOpen == false {
					result = result + "<strong>"
					strongTagOpen = true
				}

				if emTagOpen == false {
					result = result + "<em>"
					emTagOpen = true
				}

				// ....WORD***^
			} else if lastChar != char && nextChar == char && nextNextChar == char && unicode.IsLetter(nextNextNextCharRune) == false {
				if emTagOpen == true {
					result = result + "</em>"
					emTagOpen = false
				}

				if strongTagOpen == true {
					result = result + "</strong>"
					strongTagOpen = false
				}

				// ^**WORD....
			} else if unicode.IsLetter(lastCharRune) == false && nextChar == char && nextNextChar != char {
				if strongTagOpen == false {
					result = result + "<strong>"
					strongTagOpen = true
				}

				// ....WORD**^
			} else if lastChar != char && nextChar == char && unicode.IsLetter(nextNextCharRune) == false {
				if strongTagOpen == true {
					result = result + "</strong>"
					strongTagOpen = false
				}

				// ^*WORD....
			} else if unicode.IsLetter(lastCharRune) == false && nextChar != char {
				if emTagOpen == false {
					result = result + "<em>"
					emTagOpen = true
				}

				// ....WORD*^
			} else if lastChar != char && unicode.IsLetter(nextCharRune) == false {
				if emTagOpen == true {
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

				// ^~~WORD....
			} else if unicode.IsLetter(lastCharRune) == false && nextChar == "~" && nextNextChar != "~" {
				if delTagOpen == false {
					result = result + "<del>"
					delTagOpen = true
				}

				// ....WORD~~^
			} else if lastChar != "~" && nextChar == "~" && unicode.IsLetter(nextNextCharRune) == false {
				if delTagOpen == true {
					result = result + "</del>"
					delTagOpen = false
				}
			}

			// ^ - space
			// Code quote
		} else if char == "`" {
			// WORD`WORD
			if lastChar != " " && nextChar != " " {
				result = result + "`"

				// a`a
			} else if unicode.IsLetter(lastCharRune) && unicode.IsLetter(nextCharRune) {
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
