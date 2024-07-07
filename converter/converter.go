package converter

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
)

type node struct {
	__type      string
	tag         string
	content     string
	selfClosing bool
	children    []*node
}

func pop(s *[]*node) {
	if len(*s) == 0 {
		return
	}

	*s = (*s)[:len(*s)-1]
}

func getParentNode(s *[]*node) *node {
	return (*s)[len(*s)-1]
}

func generateDomTree(md string) (*node, error) {

	root := &node{__type: "container", tag: "div", content: "", children: nil}

	parentTrackStack := []*node{root}
	if md == "" {
		return nil, errors.New("string cannot be empty")
	}

	headerPattern := regexp.MustCompile(`^(#{1,6})\s*(.*)`)
	listItemPattern := regexp.MustCompile(`^[-*]\s+(.*)`)
	hrPattern := regexp.MustCompile(`^[-]{3,}$`)

	scanner := bufio.NewScanner(strings.NewReader(md))

	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(strings.TrimSpace(line), " ")

		switch {
		case headerPattern.MatchString(strings.TrimSpace(line)):
			if parentTrackStack[len(parentTrackStack)-1].tag == "ul" {
				pop(&parentTrackStack)
			}
			headerTag := &node{__type: "text", tag: "h1", content: strings.Join(arr[1:], " "), children: nil}
			parent := getParentNode(&parentTrackStack)
			parent.children = append(parent.children, headerTag)

		case listItemPattern.MatchString(strings.TrimSpace(line)):
			parent := getParentNode(&parentTrackStack)
			xr := &node{__type: "list", tag: "li", content: strings.Join(arr[1:], " "), children: nil}
			if parent.tag == "ul" {
				parent.children = append(parent.children, xr)
			} else {
				test := &node{__type: "list", tag: "ul", content: "", children: []*node{xr}}
				parent.children = append(parent.children, test)
				parentTrackStack = append(parentTrackStack, test)

			}

		case hrPattern.MatchString(strings.TrimSpace(line)):
			if getParentNode(&parentTrackStack).tag == "ul" {
				pop(&parentTrackStack)
			}
			xr := &node{__type: "line", tag: "hr", content: "", children: nil, selfClosing: true}
			parent := getParentNode(&parentTrackStack)
			parent.children = append(parent.children, xr)

		default:

		}
	}

	return root, nil
}

func (n *node) traverseDomTree() string {
	if n == nil {
		return ""
	}

	if n.selfClosing {
		return fmt.Sprintf("<%v/>", n.tag)
	}

	s := fmt.Sprintf("<%v>", n.tag)
	if n.content != "" {
		s += fmt.Sprintf(" %v", n.content)
	}
	for _, val := range n.children {
		s += val.traverseDomTree()
	}
	s += fmt.Sprintf("</%v>", n.tag)
	return s
}

func ConvertMdToHTML(text string) string {

	node, error := generateDomTree(text)
	fmt.Printf("node: %v\n", node)

	if error != nil {
		log.Fatal(error)
	}
	html := node.traverseDomTree()

	return html
}
