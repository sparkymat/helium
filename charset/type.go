package charset

// A Type specifies the type of charset
type Type string

// UTF_8 represents the UTF-8 charset
const UTF_8 Type = "UTF-8"

// ISO_8859_1 represents the ISO-8859-1 charset
const ISO_8859_1 Type = "ISO-8859-1"

func (t Type) String() string {
	return string(t)
}
