package markdown

import (
	"github.com/ckeyer/blackfriday"
)

const (
	htmlFlags = 0 |
		blackfriday.HTML_USE_XHTML |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_SMARTYPANTS_FRACTIONS |
		blackfriday.HTML_SMARTYPANTS_LATEX_DASHES

	extensions = 0 |
		blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_HEADER_IDS |
		blackfriday.EXTENSION_BACKSLASH_LINE_BREAK |
		blackfriday.EXTENSION_DEFINITION_LISTS |
		blackfriday.EXTENSION_HARD_LINE_BREAK
)

// 将markdown文本转换为html
func Trans2html(markdown []byte) (html []byte) {
	renderer := blackfriday.HtmlRenderer(htmlFlags, "", "")
	outputOpts := blackfriday.Options{
		Extensions: extensions,
	}
	// Render the markdown file into HTML and return a new []Byte
	html = blackfriday.MarkdownOptions(markdown, renderer, outputOpts)
	return
}
