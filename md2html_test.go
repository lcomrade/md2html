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
	"testing"
)

type testDataType struct {
	Input        string
	ExpectResult string
}

func TestConvert(t *testing.T) {
	// Test data
	testData := []testDataType{
		// Headers
		{
			Input:        "# My Header",
			ExpectResult: "<h1 id='my-header'>My Header</h1>",
		},
		{
			Input:        "## My Header",
			ExpectResult: "<h2 id='my-header'>My Header</h2>",
		},
		{
			Input:        "### My Header",
			ExpectResult: "<h3 id='my-header'>My Header</h3>",
		},
		{
			Input:        "#### My Header",
			ExpectResult: "<h4 id='my-header'>My Header</h4>",
		},
		{
			Input:        "##### My Header",
			ExpectResult: "<h5 id='my-header'>My Header</h5>",
		},
		{
			Input:        "###### My Header",
			ExpectResult: "<h6 id='my-header'>My Header</h6>",
		},
		{
			Input:        "######## My Header",
			ExpectResult: "<p>######## My Header</p>",
		},
		{
			Input:        "##My Header",
			ExpectResult: "<p>##My Header</p>",
		},
		{
			Input:        "### My      Header",
			ExpectResult: "<h3 id='my-header'>My      Header</h3>",
		},
		{
			Input:        "### My Header",
			ExpectResult: "<h3 id='my-header'>My Header</h3>",
		},
		{
			Input:        "### My Header #######",
			ExpectResult: "<h3 id='my-header'>My Header</h3>",
		},
		{
			Input:        "# My **Header**",
			ExpectResult: "<h1 id='my-header'>My <strong>Header</strong></h1>",
		},
		{
			Input:        "# My `good code` Header",
			ExpectResult: "<h1 id='my-good-code-header'>My <code>good code</code> Header</h1>",
		},
		// Text style
		{
			Input:        "*test test*",
			ExpectResult: "<p><em>test test</em></p>",
		},
		{
			Input:        "**test test**",
			ExpectResult: "<p><strong>test test</strong></p>",
		},
		{
			Input:        "***test test***",
			ExpectResult: "<p><strong><em>test test</em></strong></p>",
		},
		{
			Input:        "_test test_",
			ExpectResult: "<p><em>test test</em></p>",
		},
		{
			Input:        "__test test__",
			ExpectResult: "<p><strong>test test</strong></p>",
		},
		{
			Input:        "___test test___",
			ExpectResult: "<p><strong><em>test test</em></strong></p>",
		},
		{
			Input:        "~~test test~~",
			ExpectResult: "<p><del>test test</del></p>",
		},
		{
			Input:        "test * test *",
			ExpectResult: "<p>test * test *</p>",
		},
		{
			Input:        " aa*aa",
			ExpectResult: "<p> aa*aa</p>",
		},
		// Link and image
		{
			Input:        "[example domain](https://example.org)",
			ExpectResult: "<p><a href='https://example.org'>example domain</a></p>",
		},
		{
			Input:        "You can visit [example domain](https://example.org).",
			ExpectResult: "<p>You can visit <a href='https://example.org'>example domain</a>.</p>",
		},
		{
			Input:        "![Alt text](https://example.org/image.png)",
			ExpectResult: "<p><img src='https://example.org/image.png' alt='Alt text'></p>",
		},
		// Autolink
		{
			Input:        "user@example.org",
			ExpectResult: "<p><a href='mailto:user@example.org'>user@example.org</a></p>",
		},
		{
			Input:        "My email <user@example.org>.",
			ExpectResult: "<p>My email &lt<a href='mailto:user@example.org'>user@example.org</a>&gt.</p>",
		},
		{
			Input:        "@user",
			ExpectResult: "<p>@user</p>",
		},
		{
			Input:        "user@",
			ExpectResult: "<p>user@</p>",
		},
		{
			Input:        "http://example.org",
			ExpectResult: "<p><a href='http://example.org'>http://example.org</a></p>",
		},
		{
			Input:        "Semantic versioning is used (https://semver.org/).",
			ExpectResult: "<p>Semantic versioning is used (<a href='https://semver.org/'>https://semver.org/</a>).</p>",
		},
		// Code
		{
			Input:        "`man whereis`",
			ExpectResult: "<p><code>man whereis</code></p>",
		},
		{
			Input:        "aaa`aaa",
			ExpectResult: "<p>aaa`aaa</p>",
		},
		// Unordered list
		{
			Input: `
- level 1
- level 1
    - level 2
    - level 2
        - level 3
`,
			ExpectResult: "<ul><li>level 1</li><li>level 1</li><ul><li>level 2</li><li>level 2</li><ul><li>level 3</li></ul></ul></ul>",
		},
		{
			Input: `
+ level 1
+ level 1
    + level 2
    + level 2
`,
			ExpectResult: "<ul><li>level 1</li><li>level 1</li><ul><li>level 2</li><li>level 2</li></ul></ul>",
		},
		{
			Input: `
* level 1
* level 1
    * level 2
    * level 2
`,
			ExpectResult: "<ul><li>level 1</li><li>level 1</li><ul><li>level 2</li><li>level 2</li></ul></ul>",
		},
		{
			Input: `
- level 1
+ level 1
    * level 2
    + level 2
`,
			ExpectResult: "<ul><li>level 1</li><li>level 1</li><ul><li>level 2</li><li>level 2</li></ul></ul>",
		},
		// Numbered list
		{
			Input: `
1. level 1
2. level 1
    1. level 2
    2. level 2
`,
			ExpectResult: "<ol><li>level 1</li><li>level 1</li><ol><li>level 2</li><li>level 2</li></ol></ol>",
		},
		{
			Input: `
0. level 1
10000. level 1
    200. level 2
    20. level 2
`,
			ExpectResult: "<ol><li>level 1</li><li>level 1</li><ol><li>level 2</li><li>level 2</li></ol></ol>",
		},
		// Character shielding
		{
			Input:        `\*\*test test\*\*`,
			ExpectResult: "<p>**test test**</p>",
		},
		{
			Input:        `\_\_test test\_\_`,
			ExpectResult: "<p>__test test__</p>",
		},
		{
			Input:        `\~\~test test\~\~`,
			ExpectResult: "<p>~~test test~~</p>",
		},
		{
			Input:        `\#\# test test`,
			ExpectResult: "<p>## test test</p>",
		},
		{
			Input:        `\### test test`,
			ExpectResult: "<p>### test test</p>",
		},
		{
			Input: `
\+ Not list!
\+ Not list!
`,
			ExpectResult: "<p>+ Not list! + Not list!</p>",
		},
		{
			Input: `
\- Not list!
\- Not list!
`,
			ExpectResult: "<p>- Not list! - Not list!</p>",
		},
		{
			Input: `
1\. Not list!
2\. Not list!
`,
			ExpectResult: "<p>1. Not list! 2. Not list!</p>",
		},
		// Paragraphs
		{
			Input: `
The weather is good in Santo Monico.
It's always cold in Alaska.
`,
			ExpectResult: "<p>The weather is good in Santo Monico. It's always cold in Alaska.</p>",
		},
		{
			Input: `
Paragraph number 1.

Paragraph number 2.
`,
			ExpectResult: "<p>Paragraph number 1.</p><p>Paragraph number 2.</p>",
		},
		// Embedded HTML
		{
			Input: `
Code: <code>my code</code>
`,
			ExpectResult: "<p>Code: <code>my code</code></p>",
		},
		// Unicode
		{
			Input:        "Текст на русском языке",
			ExpectResult: "<p>Текст на русском языке</p>",
		},
		{
			Input:        "# Заголовок на русском языке",
			ExpectResult: "<h1 id='заголовок-на-русском-языке'>Заголовок на русском языке</h1>",
		},
	}

	// Run tests
	for _, test := range testData {
		result := Convert(test.Input)
		if result != test.ExpectResult {
			t.Error("\n" + "Input:    " + test.Input + "\n" + "Expected: " + test.ExpectResult + "\n" + "But get:  " + result)
		}
	}
}
