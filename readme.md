# Markdown to HTML Converter

This Go project converts a markdown file to HTML by first generating a DOM (Document Object Model) from the markdown and then converting this DOM into HTML.

## How It Works

1. **Parse Markdown and Generate DOM**: The `parser` package converts the markdown content into a tree-like DOM structure.
2. **Convert DOM to HTML**: The `html` package traverses the DOM and generates the corresponding HTML.

## Installation

   ```sh
   git clone https://github.com/yourusername/markdown-to-html.git
   cd markdown-to-html
   go build
   ./markdown-to-html example.md > output.html
   ```

### TODO
- [ ] Add support the nested links
- [ ] Add support for in-between tags like, bold, italic, etc...






