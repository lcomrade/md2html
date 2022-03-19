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
//   ## Closed H2 title #
//   ## Closed H2 title ##
//   ## Closed H2 title ######
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
//   ````markdown
//   ```
//   go doc
//   go tool dist list
//   go help build
//   ```
//   ````
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
//
// Character shielding:
//   | Input | Output |
//   |-------|--------|
//   | \*    | *      |
//   | \_    | _      |
//   | \~    | ~      |
//   | \#    | #      |
//   | \+    | +      |
//   | \-    | -      |
//   | \.    | .      |
//
// Paragraphs:
// A blank line is used to separate paragraphs.
//
// Embedded HTML:
// If the HTML is not in a code block, it will not be escaped.
// That mean, you can use HTML tags inside a Markdown document.
func Convert(text string) string {
	var result string = ""

	// Split text into lines
	lines := strings.Split(text, lineSeparator)

	// Track opened HTML tags
	var pTagInBuffer bool = false
	var codeTagOpen bool = false
	var codeLang string = ""
	var codeTagCloseLine string = ""
	var ulTagOpen int = 0
	var olTagOpen int = 0

	var buffer string = ""

	// Reading text by line
	for i := range lines {
		line := lines[i]

		// Inside <pre><code>....</code></pre>
		if codeTagOpen == true {
			// Close code block
			if line == codeTagCloseLine {
				// Highlight
				tmp, err := highlight.ByName(buffer, codeLang)
				if err == nil {
					buffer = tmp
				}

				// Save
				line = "<pre><code>\n" + buffer + "</code></pre>"
				buffer = ""
				codeLang = ""
				codeTagOpen = false

				// Continue read
			} else {
				buffer = buffer + line + "\n"
				line = ""
			}

			// Open code block <pre><code>
		} else if strings.HasPrefix(line, "```") {
			codeTagCloseLine, codeLang = mdCodeBlock(line)

			if pTagInBuffer == true {
				line = "<p>" + baseMdFormat(buffer) + "</p>"
				pTagInBuffer = false
			}

			line = ""
			buffer = ""
			codeTagOpen = true

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

			if pTagInBuffer == true {
				line = "<p>" + baseMdFormat(buffer) + "</p>"
				buffer = ""
				pTagInBuffer = false
			}

			// Other text
		} else {
			// List
			isUList, levelUList, resultUList := mdUList(line)
			isOList, levelOList, resultOList := mdOList(line)

			// List: <ul>
			if isUList == true {
				if pTagInBuffer == true {
					result = result + "<p>" + baseMdFormat(buffer) + "</p>"
					buffer = ""
					pTagInBuffer = false
				}

				line = baseMdFormat(resultUList)
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
				if pTagInBuffer == true {
					result = result + "<p>" + baseMdFormat(buffer) + "</p>"
					buffer = ""
					pTagInBuffer = false
				}

				line = baseMdFormat(resultOList)
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
				// Header format: <h1> - <h6>
				lineIsHeader, isHeader := mdTitle(line)

				// Paragraph or header?
				if isHeader == false {
					if buffer == "" {
						buffer = line
					} else {
						buffer = buffer + " " + line
					}

					line = ""

					pTagInBuffer = true

				} else {
					if pTagInBuffer == true {
						line = "<p>" + baseMdFormat(buffer) + "</p>"
						buffer = ""
						pTagInBuffer = false

					} else {
						line = ""
					}

					line = line + baseMdFormat(lineIsHeader)
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

	if pTagInBuffer == true {
		result = result + "<p>" + baseMdFormat(buffer) + "</p>"
		buffer = ""
		pTagInBuffer = false
	}

	if codeTagOpen == true {
		result = result + "</code></pre>"
		codeTagOpen = false
	}

	return result
}

func baseMdFormat(line string) string {
	// Links and images: <a> and <img>
	line = mdLink(line)
	line = mdAutolink(line)

	// Text format: <em>, <strong> and <code>
	line = mdStyle(line)

	return line
}
