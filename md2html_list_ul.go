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

func mdUList(line string) (bool, int, string) {
	var isList bool = false
	var result string = ""

	lineLen := len(line)

	var spaceNum int = 0
	var step string = "space" // space, skip, content

	for i := range line {
		char := string(line[i])
		nextChar := " "

		// Get next char
		if lineLen > i+1 {
			nextChar = string(line[i+1])
		}

		// 'space' step: spaces at start of line
		if step == "space" {
			if char == " " {
				spaceNum = spaceNum + 1

			} else if char == "*" && nextChar == " " {
				step = "skip"

			} else if char == "+" && nextChar == " " {
				step = "skip"

			} else if char == "-" && nextChar == " " {
				step = "skip"

			} else {
				return false, 0, line
			}

			// 'skip' step: skip *, +, -
		} else if step == "skip" {
			step = "content"
			isList = true

			// 'content' step: read list item
		} else if step == "content" {
			result = result + char
		}
	}

	// Get list level
	var level int = 1

	for spaceNum >= level*4 {
		level = level + 1
	}

	return isList, level, result
}
