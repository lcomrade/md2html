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

package md2html_test

import (
	"fmt"
	"github.com/wallblog/md2html"
)

func ExampleConvert() {
	result := md2html.Convert(`
# Header H1
## Header H2
### Header H3
#### Header H4
##### Header H5
###### Header H6

- *Italic* and _Italic_
- **Bold** and __Bold__
- ~~Strikethrough~~

[Example link](https://example.org)

1. level 1
2. level 1
    1. level 2
    2. level 2
`)

	fmt.Println(result)
}
