# Changelog
Semantic versioning is used (https://semver.org/).

## v2.0.1
- Fix: \`\`\` in row
- Fix: `*`, `-`, `~` inside word
- Fix: strong and em tag close
- Fix: `<>` brackets in non links

## v2.0.0
There is nothing dramatically new in this release.
But the quality of Markdown detection has improved noticeably.

- Closing headers are now supported
- Allow escape backtick, `+`, `-` and `#`
- Cleaning the HTML from extra spaces
- Add support opening and closing code block with 4 reverse apostrophes
- Fix: HTML ID generation
- Fix: image
- Fix: `))` in links
- Fix: email detection
- Fix: remove brackets around URL

## v1.2.0
- Autolink now work in lists
- Autolink now work in headers
- Update link parser
- Update HTML headers IDs

## v1.1.2
Fix: `mailto:'admin@example.org'` replaced to `mailto:admin@example.org`

## v1.1.1
Fix: autolink email display.

## v1.1.0
Now automatically converted into links:
- Emails
- `http://*`
- `https://*`
- `ftp://*`
- `irc://*`

## v1.0.3
The end of the list is now indicated by an empty line.

## v1.0.2
Fix: `_` and `*` inside words are now handled correctly.

## v1.0.1
Solved the problem with the display of characters of alphabets other than English.
For example, now displays Cyrillic correctly.

## v1.0.0
First stable release.

The following Markdown elements are supported:
- Paragraph
- Header
- Text style
- Link
- Image
- Code quote
- Unordered list
- Numbered list
- Character shielding(`\*`, `\_`, `\~`, `\.`)
- Embedded HTML tags

And there is quite detailed documentation and examples.
