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

// Convert http://*, https://*,
// ftp://*, irc://*
// and email to links
func mdAutolink(line string) string {
	var result string = ""

	lineSplit := strings.Split(line, " ")
	lineLen := len(lineSplit)

	for i := range lineSplit {
		buffer := lineSplit[i]

		// http://*
		if strings.HasPrefix(buffer, "http://") {
			buffer = "<a href='" + buffer + "'>" + buffer + "</a>"

			// https://*
		} else if strings.HasPrefix(buffer, "https://") {
			buffer = "<a href='" + buffer + "'>" + buffer + "</a>"

			// ftp://*
		} else if strings.HasPrefix(buffer, "ftp://") {
			buffer = "<a href='" + buffer + "'>" + buffer + "</a>"

			// irc://*
		} else if strings.HasPrefix(buffer, "irc://") {
			buffer = "<a href='" + buffer + "'>" + buffer + "</a>"

			// Email
		} else if isEmail(buffer) {
			buffer = "<a href='mailto:" + buffer + "'>" + buffer + "</a>"
		}

		result = result + buffer

		if lineLen != i+1 {
			result = result + " "
		}
	}

	return result
}

// Checks if the string is an email address.
// More about email addres format in Wikipedia:
// https://en.wikipedia.org/wiki/Email_address
func isEmail(line string) bool {
	lineRune := []rune(line)
	lineLen := len(lineRune)

	var signChar bool = false // '@' char

	lastChar := " "

	for i := range lineRune {
		charRune := lineRune[i]
		char := string(charRune)

		// Sign
		if char == "@" {
			// '@' repeated two
			if signChar == true {
				return false
			}

			// '@' is at beginning of line
			if i == 0 {
				return false
			}

			// '.@'
			if lastChar == "." {
				return false
			}

			signChar = true

			// Dot
		} else if char == "." {
			// '.abcd@example.org'
			if i == 0 {
				return false
			}

			// 'abcd@.example.org'
			if lastChar == "@" {
				return false
			}

			// 'abcd@example.org.'
			if i == lineLen-1 {
				return false
			}

			// 'ab..cd@example.org' or 'abcd@exa..mple.org'
			if lastChar == "." {
				return false
			}

			// Numbers
		} else if isNumber(charRune) {
			//pass

			// Latin letters
		} else if isLatinLetter(charRune) {
			//pass

			// Printable chars
		} else if isPrintableChar(charRune) {
			//pass

			// Unknown char
		} else {
			return false
		}

		lastChar = char
	}

	// If '@' char not exist in line
	if signChar == true {
		return true
	}

	return false
}
