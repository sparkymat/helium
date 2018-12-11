package helium

// Head represents the <head> section of an HTML document
type Head struct {
}

func (h Head) String() string {
	headString := `  <head>
  </head>`

	return headString
}
