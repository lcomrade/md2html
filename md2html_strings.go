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

// If line consists of one repeated character
func isOneCharLine(line string, char string) bool {
	if line == "" && char != "" {
		return false
	}

	lineRune := []rune(line)

	for i := range lineRune {
		if string(lineRune[i]) != char {
			return false
		}
	}

	return true
}

// Checks if character is number
func isNum(char string) bool {
	switch char {
	case "0":
		return true
	case "1":
		return true
	case "2":
		return true
	case "3":
		return true
	case "4":
		return true
	case "5":
		return true
	case "6":
		return true
	case "7":
		return true
	case "8":
		return true
	case "9":
		return true
	}

	return false
}
