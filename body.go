package helium

import (
	"fmt"
	"strings"
)

// Body represents the <body> section of an HTML document
type Body struct {
	Nodes []Node
}

// Render returns the HTML of the <body> node
func (b Body) Render() string {
	var nodesStringBuilder strings.Builder

	for _, n := range b.Nodes {
		nodesStringBuilder.WriteString("\n")
		nodesStringBuilder.WriteString(n.Render())
	}
	bodyString := fmt.Sprintf(`<body>%v
</body>`, nodesStringBuilder.String())

	return bodyString
}
