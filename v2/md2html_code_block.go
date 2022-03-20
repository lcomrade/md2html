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
	"github.com/lcomrade/highlight"
)

// Parse the line that opens the code block.
// Return open line (e.g ``` or `````) and language name (e.g. "" or "python").
//
// WARNING: You must be sure that the line is correct.
func mdCodeBlock(line string) (string, string) {
	openLine := ""
	codeLang := ""

	lineRune := []rune(line)

	for i := range lineRune {
		charRune := lineRune[i]

		if charRune == '`' {
			openLine = openLine + "`"

		} else {
			codeLang = codeLang + string(charRune)
		}
	}

	return openLine, codeLang
}

// Trying to highlight code syntax.
func tryHighlight(code string, language string) string {
	result, err := highlight.ByName(code, language)
	if err != nil {
		return code
	}

	return result
}
