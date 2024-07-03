package converter

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type node struct {
	_type   string
	content string
	tag     string
}

type branch struct {
	branchNode []node
	__type     string
}

type domTree struct {
	nodes []branch
}

func generateDomTree(md string, dom *domTree) {

	scanner := bufio.NewScanner(strings.NewReader(md))
	i := 0

	for scanner.Scan() {
		line := scanner.Text()
		allString := strings.Split(strings.TrimSpace(line), " ")

		headerPattern := regexp.MustCompile(`^(#{1,6})\s*(.*)`)
		listItemPattern := regexp.MustCompile(`^[-*]\s+(.*)`)
		// hrPattern := regexp.MustCompile(`^[-]{3,}$`)
		// boldPattern := regexp.MustCompile(`\*\*(.*?)\*\*`)
		// italicPattern := regexp.MustCompile(`_(.*?)_`)
		// blockquotePattern := regexp.MustCompile(`^>\s+(.*)`)
		// codeBlockPattern := regexp.MustCompile("```")

		switch {
		case headerPattern.MatchString(strings.TrimSpace(line)):
			branchNode := node{
				tag:     "h" + strconv.Itoa(len(allString[0])),
				content: strings.Join(allString[1:], " "),
			}

			branch := branch{branchNode: []node{branchNode}, __type: "header"}
			dom.nodes = append(dom.nodes, branch)
			i++

		case listItemPattern.MatchString(strings.TrimSpace(line)):
			branchNode := node{
				tag:     "li",
				content: strings.Join(allString[1:], " "),
				_type:   "list",
			}

			if dom.nodes[i-1].__type == "list" {
				dom.nodes[i-1].branchNode = append(dom.nodes[i-1].branchNode, branchNode)
			} else {
				branch := branch{branchNode: []node{branchNode}, __type: "list"}
				dom.nodes = append(dom.nodes, branch)
				i++

			}

		}

	}

}

func ConvertToHTML(input string) (output string) {
	dom := &domTree{nodes: []branch{}}
	html := ""

	generateDomTree(input, dom)

	fmt.Printf("%+v", dom)

	for _, value := range dom.nodes {
		listCounter := len(value.branchNode)
		for _, val := range value.branchNode {
			switch value.__type {
			case "header":
				html += fmt.Sprintf("<%v>%v</%v>\n", val.tag, val.content, val.tag)

			case "list":
				if listCounter == len(value.branchNode) {
					html += "<ul>\n"
				}
				listCounter--

				html += fmt.Sprintf("<%v>%v</%v>\n", val.tag, val.content, val.tag)
				if listCounter <= 0 {
					html += "</ul>\n"
				}

			}
		}

	}

	return html
}
