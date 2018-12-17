package helium

import (
	"fmt"
	"strings"

	"github.com/sparkymat/helium/charset"
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

// MetaCharset represents a <meta> tag for setting the charset
func MetaCharset(value charset.Type) HeadNode {
	node := HeadNode{name: "meta", hideClosingTag: true}
	node.Attr("charset", value.String())
	return node
}
