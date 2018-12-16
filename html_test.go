package helium_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/sparkymat/helium"
)

var _ = Describe("Html", func() {
	It("returns a node representing the html document", func() {
		doc := HTML{}

		expectedString := `<!DOCTYPE html>
<html language="en">
<head>
<meta charset="utf-8">
</head>
<body>
</body>
</html>
`

		Expect(doc.String()).To(Equal(expectedString))
	})
})
