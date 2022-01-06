[![Go report](https://goreportcard.com/badge/github.com/wallblog/md2html)](https://goreportcard.com/report/github.com/wallblog/md2html)
[![Go Reference](https://pkg.go.dev/badge/github.com/wallblog/md2html.svg)](https://pkg.go.dev/github.com/wallblog/md2html)
[![Release](https://img.shields.io/github/v/release/wallblog/md2html)](https://github.com/wallblog/md2html/releases/latest)
[![License](https://img.shields.io/github/license/wallblog/md2html)](LICENSE)

**md2html** is a golang library for converting Markdown to HTML.

## Installation
```
go get github.com/wallblog/md2html
```

## Using
```go
package main

import(
	"github.com/wallblog/md2html"
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
`

func main() {
	result := md2html.Convert(myMarkdown)

	println(result)
}
```
