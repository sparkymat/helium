package helium

import (
	"fmt"
	"strings"
)

// Body represents the <body> section of an HTML document
type body struct {
	nodes []Node
}

func Body(children ...Node) body {
	return body{children}
}

// Render returns the HTML of the <body> node
func (b body) Render() string {
	var nodesStringBuilder strings.Builder

	for _, n := range b.nodes {
		nodesStringBuilder.WriteString("\n")
		nodesStringBuilder.WriteString(n.Render())
	}
	bodyString := fmt.Sprintf(`<body>%v
</body>`, nodesStringBuilder.String())

	return bodyString
}
