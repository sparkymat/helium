package helium

import "fmt"

// HTML represents a valid HTML document
type HTML struct {
	Head head
	Body body
}

func (h HTML) Render() string {
	htmlString := `<!DOCTYPE html>
<html language="en">
%v
%v
</html>
`

	return fmt.Sprintf(htmlString, h.Head.Render(), h.Body.Render())
}
