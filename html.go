package helium

import "fmt"

// HTML represents a valid HTML document
type HTML struct {
	Head Head
	Body Body
}

func (h HTML) String() string {
	htmlString := `<!DOCTYPE html>
<html language="en">
%v
%v
</html>
`

	return fmt.Sprintf(htmlString, h.Head.String(), h.Body.String())
}
