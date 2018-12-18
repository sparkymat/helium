package helium_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/sparkymat/helium"
	"github.com/sparkymat/helium/charset"
)

var _ = Describe("Html", func() {
	It("returns a node representing the html document", func() {
		doc := HTML{
			Head(
				MetaCharset(charset.UTF_8),
				MetaViewportDefault(),
				Title("Hello World"),
			),
			Body(),
		}

		expectedString := `<!DOCTYPE html>
<html language="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>Hello World</title>
</head>
<body>
</body>
</html>
`

		Expect(doc.Render()).To(Equal(expectedString))
	})
})
