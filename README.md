[![Go report](https://goreportcard.com/badge/github.com/lcomrade/md2html)](https://goreportcard.com/report/github.com/lcomrade/md2html)
[![Go Reference](https://pkg.go.dev/badge/github.com/lcomrade/md2html.svg)](https://pkg.go.dev/github.com/lcomrade/md2html#section-documentation)
[![Release](https://img.shields.io/github/v/release/lcomrade/md2html)](https://github.com/lcomrade/md2html/releases/latest)
[![License](https://img.shields.io/github/license/lcomrade/md2html)](LICENSE)

**md2html** is a golang library for converting Markdown to HTML.

## Install
Supported Go versions:
- 1.11
- 1.17
- 1.18

Add to `go.mod` file:
```go.mod
require github.com/lcomrade/md2html/v2 v2
```

## Example
```go
package main

import(
	"github.com/lcomrade/md2html/v2"
)

const myMarkdown = `
# Title
Some text here.

*Italic*
**Bold**
~~Strikethrough~~

1. level 1
2. level 1
3. level 1
    1. level 2
    2. level 2
        1. level 3
        2. level 3
`

func main() {
	result := md2html.Convert(myMarkdown)
	println(result)
}
```

## Documentation
- Offline documentation: `go doc -all github.com/lcomrade/md2html`
- [Online documentation](https://pkg.go.dev/github.com/lcomrade/md2html#section-documentation)
- [Markdown Syntax Guide](docs/syntax_guide.md)
- [Changelog](CHANGELOG.md)
