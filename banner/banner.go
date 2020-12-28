// Package banner contains general utilities to display a banner
package banner

import (
	"fmt"
	"strings"
)

// Banner The Banner layout properties
type Banner struct {
	TlChar     rune // top left char
	TrChar     rune // top right char
	BlChar     rune // bottom left char
	BrChar     rune // bottom right char
	DashesChar rune
	SidesChar  rune
	Message    string
	Padding    uint8
}

// NewBanner Creates a new banner instance
func NewBanner(tlChar, trChar, blChar, brChar, dashesChar, sidesChar rune, message string, padding uint8) Banner {
	return Banner{
		TlChar:     tlChar,
		TrChar:     trChar,
		BlChar:     blChar,
		BrChar:     brChar,
		DashesChar: dashesChar,
		SidesChar:  sidesChar,
		Message:    message,
		Padding:    padding,
	}
}

//NewDefaultBanner Creates a banner with the default options
func NewDefaultBanner(msg string) Banner {
	return NewBanner('+', '*', '*', '+', '-', '|', msg, 2)
}

// PrintHorizontalEdges Prints the banner's horizontal edges
func (b Banner) PrintHorizontalEdges() string {
	msgLen := len(b.Message)
	result := fmt.Sprintf("%c%s%c",
		b.TlChar,
		strings.Repeat(string(b.DashesChar), msgLen+int(b.Padding)*2),
		b.TrChar)

	return result
}

// PrintMiddle Prints the middle of the banner
func (b Banner) PrintMiddle() string {
	paddingsRep := strings.Repeat(" ", int(b.Padding))
	return fmt.Sprintf("%c%s%s%s%c", b.SidesChar, paddingsRep, b.Message, paddingsRep, b.SidesChar)
}
