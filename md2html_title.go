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

// Replace # - ###### to <h1> - <h6>
// Return formated string and 'is header' boolean
func mdTitle(line string) (string, bool) {
	lineRune := []rune(line)

	if strings.HasPrefix(line, "# ") {
		title := string(lineRune[2:])
		title = trimTitleSharp(title)
		return "<h1 id='" + toHTMLID(title) + "'>" + title + "</h1>", true
	}

	if strings.HasPrefix(line, "## ") {
		title := string(lineRune[3:])
		title = trimTitleSharp(title)
		return "<h2 id='" + toHTMLID(title) + "'>" + title + "</h2>", true
	}

	if strings.HasPrefix(line, "### ") {
		title := string(lineRune[4:])
		title = trimTitleSharp(title)
		return "<h3 id='" + toHTMLID(title) + "'>" + title + "</h3>", true
	}

	if strings.HasPrefix(line, "#### ") {
		title := string(lineRune[5:])
		title = trimTitleSharp(title)
		return "<h4 id='" + toHTMLID(title) + "'>" + title + "</h4>", true
	}

	if strings.HasPrefix(line, "##### ") {
		title := string(lineRune[6:])
		title = trimTitleSharp(title)
		return "<h5 id='" + toHTMLID(title) + "'>" + title + "</h5>", true
	}

	if strings.HasPrefix(line, "###### ") {
		title := string(lineRune[7:])
		title = trimTitleSharp(title)
		return "<h6 id='" + toHTMLID(title) + "'>" + title + "</h6>", true
	}

	return line, false
}

// Convert '## Some title ####' to '## Some title'
func trimTitleSharp(line string) string {
	lineRune := []rune(line)
	lineLen := len(lineRune)

	endI := 0

	for i := range line {
		char := lineRune[lineLen-i-1]

		if char != '#' {
			break
		}

		endI = endI + 1
	}

	if endI != 0 {
		return string(line[:lineLen-endI-1])
	}

	return line
}
