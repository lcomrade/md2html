# Changelog
Semantic versioning is used (https://semver.org/).

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
