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

// Convert http:// and https:// to links
func mdAutolink(line string) string {
	var result string = ""

	lineSplit := strings.Split(line, " ")
	lineLen := len(lineSplit)

	for i := range lineSplit {
		buffer := lineSplit[i]
		bufferLen := len(buffer)

		// http://*
		if bufferLen > 7 && strings.HasPrefix(buffer, "http://") {
			buffer = "<a href='" + buffer + "'>" + buffer + "</a>"

			// https://*
		} else if bufferLen > 8 && strings.HasPrefix(buffer, "https://") {
			buffer = "<a href='" + buffer + "'>" + buffer + "</a>"
		}

		result = result + buffer

		if lineLen != i+1 {
			result = result + " "
		}
	}

	return result
}
