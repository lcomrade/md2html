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
// Headers:
//   # This is a H1 header
//   ## This is a H2 header
//   ### This is a H3 header
//   #### This is a H4 header
//   ##### This is a H5 header
//   ###### This is a H6 header
//
// Text style ('*' may be replaced by '_'):
//   *Italic text*
//   **Bold text**
//   ***Bold and italic text***
//   ~~Strikethrough text~~
//
// Link and image:
//   [Example link](https://example.org)
//   ![Alt text](https://example.org/image.png)
//
// Autolink:
//   admin@example.org
//   http://example.org
//   https://example.org
//   ftp://ftp.mozilla.org
//   irc://irc.debian.org/debian
//
// Code:
//   `Code quote`
//
//   ```
//   go doc
//   go tool dist list
//   go help build
//   ```
//
// Unordered list ('-' may be replaced by '+' or '*'):
//   - level 1
//   - level 1
//       - level 2
//       - level 2
//   ^^^^
//   (4 spaces)
//
// Numbered list:
//   1. level 1
//   2. level 1
//       1. level 2
//       2. level 2
//   ^^^^
//   (4 spaces)
func Convert(text string) string {
	var result string = ""

	// Split text into lines
	lines := strings.Split(text, lineSeparator)

	// Track opened HTML tags
	var pTagOpen bool = false
	var codeTagOpen bool = false
	var ulTagOpen int = 0
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
			// Close <ul>, <ol> and <p> tags
			for ulTagOpen != 0 {
				result = result + "</ul>"
				ulTagOpen = ulTagOpen - 1
			}

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
			// List
			isUList, levelUList, resultUList := mdUList(line)
			isOList, levelOList, resultOList := mdOList(line)

			// List: <ul>
			if isUList == true {
				line = mdStyle(resultUList)
				line = mdLink(line)
				line = mdAutolink(line)
				line = "<li>" + line + "</li>"

				for olTagOpen != 0 {
					result = result + "</ol>"
					olTagOpen = olTagOpen - 1
				}

				for ulTagOpen < levelUList {
					line = "<ul>" + line
					ulTagOpen = ulTagOpen + 1
				}

				for ulTagOpen > levelUList {
					line = "</ul>" + line
					ulTagOpen = ulTagOpen - 1
				}

				// List: <ol>
			} else if isOList == true {
				line = mdStyle(resultOList)
				line = mdLink(line)
				line = mdAutolink(line)
				line = "<li>" + line + "</li>"

				for ulTagOpen != 0 {
					result = result + "</ul>"
					ulTagOpen = ulTagOpen - 1
				}

				for olTagOpen < levelOList {
					line = "<ol>" + line
					olTagOpen = olTagOpen + 1
				}

				for olTagOpen > levelOList {
					line = "</ol>" + line
					olTagOpen = olTagOpen - 1
				}

			} else {
				// Text format: <em>, <strong> and <code>
				line = mdStyle(line)

				// Links and images: <a> and <img>
				line = mdLink(line)
				line = mdAutolink(line)

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
	for ulTagOpen != 0 {
		result = result + "</ul>"
		ulTagOpen = ulTagOpen - 1
	}

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
