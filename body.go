package helium

// Body represents the <body> section of an HTML document
type Body struct {
}

func (b Body) Render() string {
	bodyString := `<body>
</body>`

	return bodyString
}
