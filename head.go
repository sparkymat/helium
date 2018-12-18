package helium

import (
	"fmt"
	"strings"
)

type head struct {
	nodes []HeadNode
}

// Head represents the <head> section of an HTML document
func Head(children ...HeadNode) head {
	h := head{children}
	return h
}

// Render returns the HTML of the <head> node
func (h head) Render() string {
	var nodesStringBuilder strings.Builder

	for _, n := range h.nodes {
		nodesStringBuilder.WriteString("\n")
		nodesStringBuilder.WriteString(n.Render())
	}

	headString := fmt.Sprintf(`<head>%v
</head>`, nodesStringBuilder.String())

	return headString
}

// Title represents the title for an HTML document
func Title(value string) HeadNode {
	return HeadNode{name: "title", text: value}
}
