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

// Replace [Example domain](https://example.org) to <a>
// Replace ![Alt text](https://example.org/image.png) to <img>
func mdLink(line string) string {
	var result string = ""

	lineRune := []rune(line)
	lineLen := len(lineRune)

	var nowRead string = "normal" // normal, arg1 or arg2
	var contType string = ""      // content type (link or image)
	var arg1 string = ""          // square brackets
	var arg2 string = ""          // round brackets

	for i := range lineRune {
		lastLastChar := " "
		lastCharRune := ' '
		lastChar := " "
		char := string(lineRune[i])
		nextCharRune := ' '
		nextChar := " "

		// Get last last char
		if i > 1 {
			lastLastChar = string(lineRune[i-2])
		}

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

		// Link start: ^[....
		if unicode.IsLetter(lastCharRune) == false && char == "[" {
			nowRead = "arg1"
			contType = "link"

			// Image start: ^![....
		} else if lastLastChar == " " && lastChar == "!" && char == "[" {
			nowRead = "arg1"
			contType = "image"

			// Image start (skip !): ^![....
		} else if lastChar == " " && char == "!" && nextChar == "[" {
			//pass

			// Square brackets end: ](....
		} else if lastChar != " " && char == "]" && nextChar == "(" {
			nowRead = "arg2"

			// Round brackets start: ](....
		} else if lastChar == "]" && char == "(" {
			//pass

			// Round brackets end: ....)^ or ....).^
		} else if char == ")" && nowRead == "arg2" {
			// ....))
			if nextChar == ")" {
				arg2 = arg2 + char

				// ....)^
			} else if unicode.IsLetter(nextCharRune) == false {
				if contType == "link" {
					result = result + "<a href='" + arg2 + "'>" + arg1 + "</a>"

				} else {
					result = result + "<img src='" + arg2 + "' alt='" + arg1 + "'>"
				}

				nowRead = "normal"
				contType = ""
				arg1 = ""
				arg2 = ""

				// ....)aa
			} else {
				arg2 = arg2 + char
			}

			// Normal character
		} else {
			if nowRead == "arg1" {
				arg1 = arg1 + char

			} else if nowRead == "arg2" {
				arg2 = arg2 + char

			} else {
				result = result + char
			}
		}
	}

	return result
}
