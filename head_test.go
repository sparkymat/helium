package helium_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/sparkymat/helium"
	"github.com/sparkymat/helium/charset"
)

var _ = Describe("Head", func() {
	It("returns correct nodes for the various <meta> variants", func() {
		Expect(MetaCharset(charset.UTF_8).Render()).To(Equal("<meta charset=\"UTF-8\">"))
		Expect(MetaAuthor("Adam").Render()).To(Equal("<meta name=\"author\" content=\"Adam\">"))
		Expect(MetaDescription("Something").Render()).To(Equal("<meta name=\"description\" content=\"Something\">"))
		Expect(MetaGenerator("Frontpage").Render()).To(Equal("<meta name=\"generator\" content=\"Frontpage\">"))
		Expect(MetaKeywords("foo", "bar").Render()).To(Equal("<meta name=\"keywords\" content=\"foo,bar\">"))
		Expect(MetaViewportDefault().Render()).To(Equal("<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">"))
		Expect(MetaViewportSimple("500", 2.5).Render()).To(Equal("<meta name=\"viewport\" content=\"width=500, initial-scale=2.5\">"))
		Expect(MetaViewportExtended(DeviceWidth, 1.0, 1.0, 2.5, true).Render()).To(Equal("<meta name=\"viewport\" content=\"width=device-width, initial-scale=1, minimum-scale=1, maximum-scale=2.5, user-scalable=yes\">"))
	})
})
