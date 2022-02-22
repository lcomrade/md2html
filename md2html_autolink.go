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

		// Find and remove brackets
		buffer, bufferStart, bufferEnd := removeBrackets(buffer)

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

		buffer = bufferStart + buffer + bufferEnd
		result = result + buffer

		if lineLen != i+1 {
			result = result + " "
		}
	}

	return result
}

// Removes opening and closing brackets
func removeBrackets(line string) (string, string, string) {
	lineLen := len(line)
	if lineLen <= 3 {
		return line, "", ""
	}

	brackets := []string{
		"()", "<>", "[]", "{}",
	}

	firstChar := string(line[0])
	preLastChar := string(line[lineLen-2])
	lastChar := string(line[lineLen-1])

	for _, br := range brackets {
		openBracket := string(br[0])
		closeBracket := string(br[1])

		// '(....)'
		if firstChar == openBracket && lastChar == closeBracket {
			return line[1 : lineLen-1], shieldHTML(openBracket), shieldHTML(closeBracket)
		}

		// '(....).'
		if firstChar == openBracket && preLastChar == closeBracket {
			return line[1 : lineLen-2], shieldHTML(openBracket), shieldHTML(closeBracket) + lastChar
		}
	}

	return line, "", ""
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
