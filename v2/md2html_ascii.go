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

// Checks if rune is a number (0-9).
func isNumber(char rune) bool {
	numbers := []rune{
		'0', '1', '2', '3', '4',
		'5', '6', '7', '8', '9',
	}

	for _, number := range numbers {
		if char == number {
			return true
		}
	}

	return false
}

// Checks if rune is a small Latin letter (a-z).
func isSmallLatinLetter(char rune) bool {
	letters := []rune{
		'a', 'b', 'c', 'd', 'e', 'f',
		'g', 'h', 'i', 'j', 'k', 'l',
		'm', 'n', 'o', 'p', 'q', 'r',
		's', 't', 'u', 'v', 'w', 'x',
		'y', 'z',
	}

	for _, letter := range letters {
		if char == letter {
			return true
		}
	}

	return false
}

// Checks if rune is a big Latin letter (A-Z).
func isBigLatinLetter(char rune) bool {
	letters := []rune{
		'A', 'B', 'C', 'D', 'E', 'F',
		'G', 'H', 'I', 'J', 'K', 'L',
		'M', 'N', 'O', 'P', 'Q', 'R',
		'S', 'T', 'U', 'V', 'W', 'X',
		'Y', 'Z',
	}

	for _, letter := range letters {
		if char == letter {
			return true
		}
	}

	return false
}

// Checks if rune is a Latin letter (a-z and A-Z).
func isLatinLetter(char rune) bool {
	if isSmallLatinLetter(char) {
		return true
	}

	if isBigLatinLetter(char) {
		return true
	}

	return false
}

// Checks if rune is a printable char:
//   !, #, $, %, &, ~,
//   *, +, -, /, =, ?,
//   ^, _, `, {, |, },
//   '
func isPrintableChar(char rune) bool {
	printableChars := []rune{
		'!', '#', '$', '%', '&', '~',
		'*', '+', '-', '/', '=', '?',
		'^', '_', '`', '{', '|', '}',
		[]rune(`'`)[0],
	}

	for _, prChar := range printableChars {
		if char == prChar {
			return true
		}
	}

	return false
}
