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


## Numbered list
```markdown
1. level 1
2. level 1
    1. level 2
    2. level 2
^^^^
(4 spaces)
```
