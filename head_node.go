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
	attributes     map[string]string
}

// Render returns the HTML of a head node
func (n HeadNode) Render() string {
	closingTag := ""
	if !n.hideClosingTag {
		closingTag = fmt.Sprintf("</%v>", n.name)
	}

	var attributesStringBuilder strings.Builder

	for name, value := range n.attributes {
		attributesStringBuilder.WriteString(fmt.Sprintf(" %v=\"%v\"", name, value))
	}

	return fmt.Sprintf("<%v%v>%v%v", n.name, attributesStringBuilder.String(), n.text, closingTag)
}

// Attr sets the value of an attribute
func (n *HeadNode) Attr(name string, value string) {
	if n.attributes == nil {
		n.attributes = make(map[string]string)
	}

	n.attributes[name] = value
}
