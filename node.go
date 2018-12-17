package helium

// Node represents a child node of the <body> node
type Node interface {
	Render() string
}
