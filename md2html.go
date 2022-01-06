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

const lineSeparator = "\n"

// Convert Markdown to HTML.
//
// Supports following Markdown elements:
// - Headers (# - ######)
// - Italic text (*Text* or _Text_)
// - Bold text (**Text** or __Text__)
// - Strikethrough text (~~Text~~)
// - Code quote (`Code line`)
// - Code block (```Many lines of code```)
func Convert(text string) string {
	var result string = ""

	// Split text into lines
	lines := strings.Split(text, "\n")

	// Track opened HTML tags
	var pTagOpen bool = false
	var codeTagOpen bool = false
	var olTagOpen int = 0
	var isHeader bool = false

	// Reading text by line
	for i := range lines {
		line := lines[i]

		// Shield characters inside <pre><code>....</code></pre>
		if line != "```" && codeTagOpen == true {
			line = shieldHTML(line) + lineSeparator

			// End paragraph: </p> and </ol>
		} else if line == "" {
			for olTagOpen != 0 {
				result = result + "</ol>"
				olTagOpen = olTagOpen - 1
			}

			if pTagOpen == true {
				line = "</p>"
				pTagOpen = false
			}

			// If code block: <pre><code>
		} else if line == "```" {
			if codeTagOpen == false {
				if pTagOpen == true {
					line = "</p><pre><code>"
					pTagOpen = false

				} else {
					line = "<pre><code>"
				}

				codeTagOpen = true

			} else {
				line = "</code></pre>"
				codeTagOpen = false
			}

			// Other text
		} else {
			// List: <ol>
			isOList, level, resultOList := mdOList(line)
			if isOList == true {
				line = mdStyle(resultOList)
				line = "<li>" + mdLink(line) + "</li>"

				for olTagOpen < level {
					line = "<ol>" + line
					olTagOpen = olTagOpen + 1
				}

				for olTagOpen > level {
					line = "</ol>" + line
					olTagOpen = olTagOpen - 1
				}

			} else {
				// Text format: <em>, <strong> and <code>
				line = mdStyle(line)

				// Links and images: <a> and <img>
				line = mdLink(line)

				// Header format: <h1> - <h6>
				line, isHeader = mdTitle(line)

				// Paragraph or header?
				if isHeader == false {
					if pTagOpen == false {
						line = "<p>" + line
						pTagOpen = true

					} else {
						line = " " + line
					}

				} else {
					if pTagOpen == true {
						line = "</p>" + line
						pTagOpen = false
					}

					isHeader = false
				}
			}
		}

		result = result + line
	}

	// Closing not closed HTML tags
	for olTagOpen != 0 {
		result = result + "</ol>"
		olTagOpen = olTagOpen - 1
	}

	if pTagOpen == true {
		result = result + "</p>"
		pTagOpen = false
	}

	if codeTagOpen == true {
		result = result + "</code></pre>"
		codeTagOpen = false
	}

	return result
}
