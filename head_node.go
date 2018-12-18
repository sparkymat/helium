package helium

import (
	"fmt"
	"strings"
)

// HeadNode represents the child of a <head> node
type HeadNode struct {
	name           string
	text           string
	hideClosingTag bool
	attributes     []attribute
}

// Render returns the HTML of a head node
func (n HeadNode) Render() string {
	closingTag := ""
	if !n.hideClosingTag {
		closingTag = fmt.Sprintf("</%v>", n.name)
	}

	var attributesStringBuilder strings.Builder

	for _, attr := range n.attributes {
		attributesStringBuilder.WriteString(fmt.Sprintf(" %v=\"%v\"", attr.name, attr.value))
	}

	return fmt.Sprintf("<%v%v>%v%v", n.name, attributesStringBuilder.String(), n.text, closingTag)
}

// Attr sets the value of an attribute
func (n *HeadNode) Attr(name string, value string) {
	for _, attr := range n.attributes {
		if attr.name == name {
			attr.value = value
			return
		}
	}

	n.attributes = append(n.attributes, attribute{name: name, value: value})
}
