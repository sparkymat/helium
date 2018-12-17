package helium

import (
	"fmt"
	"strings"
)

// Head represents the <head> section of an HTML document
type Head struct {
	Nodes []Node
}

func (h Head) Render() string {
	var nodesStringBuilder strings.Builder

	for _, n := range h.Nodes {
		nodesStringBuilder.WriteString("\n")
		nodesStringBuilder.WriteString(n.Render())
	}

	headString := fmt.Sprintf(`<head>
<meta charset="utf-8">%v
</head>`, nodesStringBuilder.String())

	return headString
}

// Title represnts the title for an HTML document
type Title string

func (t Title) Render() string {
	return fmt.Sprintf("<title>%v</title>", string(t))
}
