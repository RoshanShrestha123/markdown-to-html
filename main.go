package main

import (
	"fmt"

	"github.com/RoshanShrestha123/markdown-to-html/converter"
)

func main() {
	fmt.Println("Convert Markdown to HTML")
	md := `
	## Header 2
	### Header 3
	#### Header 4
	##### Header 5
	###### Header 6
	- Bullet List Item 1
	- Bullet List Item 2
	# header
	 - test

	`
	htmlCode := converter.ConvertToHTML(md)
	fmt.Println(htmlCode)
}
