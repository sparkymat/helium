package helium_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/sparkymat/helium"
)

var _ = Describe("Html", func() {
	It("returns a node representing the html document", func() {
		doc := HTML{
			Head{
				[]HeadNode{
					Title("Hello World"),
				},
			},
			Body{},
		}

		expectedString := `<!DOCTYPE html>
<html language="en">
<head>
<meta charset="utf-8">
<title>Hello World</title>
</head>
<body>
</body>
</html>
`

		Expect(doc.Render()).To(Equal(expectedString))
	})
})
