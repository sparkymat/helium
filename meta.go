package helium

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sparkymat/helium/charset"
)

// DeviceWidth represnts a viewport width of 'device-width'
const DeviceWidth string = "device-width"

func metaWithAttr(name string, content string) HeadNode {
	node := HeadNode{name: "meta", hideClosingTag: true}
	node.Attr("name", name)
	node.Attr("content", content)
	return node
}

// MetaCharset represents a <meta> tag for setting the charset
func MetaCharset(value charset.Type) HeadNode {
	node := HeadNode{name: "meta", hideClosingTag: true}
	node.Attr("charset", value.String())
	return node
}

// MetaAuthor represents a <meta> tag for setting the author
func MetaAuthor(name string) HeadNode {
	return metaWithAttr("author", name)
}

// MetaDescription represents a <meta> tag for setting the description
func MetaDescription(name string) HeadNode {
	return metaWithAttr("description", name)
}

// MetaGenerator represents a <meta> tag for setting the generator
func MetaGenerator(name string) HeadNode {
	return metaWithAttr("generator", name)
}

// MetaKeywords represents a <meta> tag for setting the keywords
func MetaKeywords(keywords ...string) HeadNode {
	return metaWithAttr("keywords", strings.Join(keywords, ","))
}

// MetaViewportDefault returns a <meta> tag for viewport with device_width and scale of 1.0
func MetaViewportDefault() HeadNode {
	return MetaViewportSimple(DeviceWidth, 1.0)
}

// MetaViewportSimple represents a <meta> tag for setting the viewport configuration with just the width and initial scale
func MetaViewportSimple(width string, initialScale float64) HeadNode {
	return metaViewport(width, initialScale, nil, nil, nil)
}

// MetaViewportExtended represents a <meta> tag for setting the viewport configuration with all options
func MetaViewportExtended(width string, initialScale float64, minimumScale float64, maximumScale float64, userScalable bool) HeadNode {
	return metaViewport(width, initialScale, &minimumScale, &maximumScale, &userScalable)
}

func metaViewport(width string, initialScale float64, minimumScale *float64, maximumScale *float64, userScalable *bool) HeadNode {
	var values []string

	values = append(values, fmt.Sprintf("width=%v", width))
	values = append(values, fmt.Sprintf("initial-scale=%v", formatFloat(initialScale)))

	if minimumScale != nil {
		values = append(values, fmt.Sprintf("minimum-scale=%v", formatFloat(*minimumScale)))
	}

	if maximumScale != nil {
		values = append(values, fmt.Sprintf("maximum-scale=%v", formatFloat(*maximumScale)))
	}

	if userScalable != nil {
		if *userScalable == true {
			values = append(values, "user-scalable=yes")
		} else {
			values = append(values, "user-scalable=no")
		}
	}

	return metaWithAttr("viewport", strings.Join(values, ", "))
}

func formatFloat(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}
