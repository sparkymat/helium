package helium

// Head represents the <head> section of an HTML document
type Head struct {
	Nodes []Node
}

func (h Head) String() string {

	headString := `<head>
<meta charset="utf-8">
</head>`

	return headString
}

// Title represnts the title for an HTML document
type Title string

func (t Title) String() string {
	return string(t)
}
