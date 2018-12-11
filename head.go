package helium

// Head represents the <head> section of an HTML document
type Head struct {
}

func (h Head) String() string {

	headString := `	<head>
		<meta charset="utf-8">
	</head>`

	return headString
}
