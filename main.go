package main

import (
	"fmt"

	"github.com/RoshanShrestha123/markdown-to-html/converter"
)

func main() {
	text := `## Header32
	## Header1
	## Header54
	- test
	- test2 
	- test 3
	-----


	## Header
	## Header

	- test man
	- test man
`

	html := converter.ConvertMdToHTML(text)
	fmt.Println(html)

}
