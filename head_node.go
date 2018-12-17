package helium

import "fmt"

// Node represents the child of a <head> node
type HeadNode struct {
	name           string
	text           string
	hideClosingTag bool
}

func (n HeadNode) Render() string {
	closingTag := ""
	if !n.hideClosingTag {
		closingTag = fmt.Sprintf("</%v>", n.name)
	}
	return fmt.Sprintf("<%v>%v%v", n.name, n.text, closingTag)
}
