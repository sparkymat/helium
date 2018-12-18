package helium

import (
	"fmt"
	"strings"
)

// Head represents the <head> section of an HTML document
type Head struct {
	Nodes []HeadNode
}

// Render returns the HTML of the <head> node
func (h Head) Render() string {
	var nodesStringBuilder strings.Builder

	for _, n := range h.Nodes {
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
