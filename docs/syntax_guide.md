# Markdown Syntax Guide
## Headers
```markdown
# This is a H1 header
## This is a H2 header
### This is a H3 header
#### This is a H4 header
##### This is a H5 header
###### This is a H6 header
```

```markdown
## Closed H2 title #
## Closed H2 title ##
## Closed H2 title ######
```



## Text style
| Style           | Markdown                           | View          | HTML                                |
| --------------- | ---------------------------------- | ------------- | ----------------------------------- |
| Italic          | `*My Text*` or `_My Text_`         | *My Text*     | `<em>My Text</em>`                  |
| Bold            | `**My Text**` or `__My Text__`     | **My Text**   | `<strong>My text</strong>`          |
| Italic and Bold | `***My Text***` or `___My Text___` | ***My Text*** | `<strong><em>My Text</em></strong>` |
| Strikethrough   | `~~My Text~~`                      | ~~My Text~~   | `<del>My Text</del>`                |


## Link and image
```markdown
[Example link](https://example.org)
![Alt text](https://example.org/image.png)
```


## Autolink
Will be preformed into links:
- Emails
- `http://*`
- `https://*`
- `ftp://*`
- `irc://*`


## Code
```markdown
`Code quote`
```

````markdown 
```
go doc
go tool dist list
go help build
```
````

`````markdown
````markdown
```bash
go doc
go tool dist list
go help build
```
````
`````

## Unordered list
NOTE: `-` may be replaced by `+` or `*`

```markdown
- level 1
- level 1
    - level 2
    - level 2
^^^^
(4 spaces)
```


## Numbered list
```markdown
1. level 1
2. level 1
    1. level 2
    2. level 2
^^^^
(4 spaces)
```

## Character shielding
| Input | Output |
| ----- | ------ |
| `\*`  | `*`    |
| `\_`  | `_`    |
| `\~`  | `~`    |
| `\#`  | `#`    |
| `\+`  | `+`    |
| `\-`  | `-`    |
| `\.`  | `.`    |

You can also use this for code quote:
```markdown
\` ----> `

\`my code\` ----> `my code`
```



## Paragraphs
A blank line is used to separate paragraphs.


## Embedded HTML
If the HTML is not in a code block, it will not be escaped.
That mean, you can use HTML tags inside a Markdown document.
