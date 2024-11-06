// Package banner contains general utilities to display a banner
package banner

import (
	"fmt"
	"strings"
)

// Banner The Banner layout properties
type Banner struct {
	TlChar, // top left char
	TrChar, // top right char
	BlChar, // bottom left char
	BrChar, // bottom right char
	DashesChar,
	SidesChar,
	Message string
	Padding uint8
}

// HorLineType A banner horizontal line types
type HorLineType int

const (
	// HorLineTypeTop the top line of a banner
	HorLineTypeTop = iota
	// HorLineTypeBottom the bottom line of a banner
	HorLineTypeBottom = iota
)

// NewBanner Creates a new banner instance
func NewBanner(tlChar, trChar, blChar, brChar, dashesChar, sidesChar string, message string, padding uint8) Banner {
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

// NewDefaultBanner Creates a banner with the default options
func NewDefaultBanner(msg string) Banner {
	return NewBanner("+", "*", "*", "+", "-", "|", msg, 2)
}

// errInvalidBannerHorLineType Invalid banner horizontal type
type errInvalidBannerHorLineType int

func (e errInvalidBannerHorLineType) Error() string {
	return fmt.Sprintf("Invalid horizontal line type: %d", e)
}

// PrintHorizontalEdges Prints the banner's horizontal edges
func (b Banner) PrintHorizontalEdges(line HorLineType) (string, error) {
	msgLen := len(b.Message)
	var lcChar, rcChar string // left corner and right corner chars
	switch line {
	case HorLineTypeTop:
		lcChar = b.TlChar
		rcChar = b.TrChar
	case HorLineTypeBottom:
		lcChar = b.BlChar
		rcChar = b.BrChar
	default:
		return "", errInvalidBannerHorLineType(line)
	}

	result := fmt.Sprintf("%s%s%s",
		lcChar,
		strings.Repeat(b.DashesChar, msgLen+int(b.Padding)*2), rcChar)

	return result, nil
}

// PrintMiddle Prints the middle of the banner
func (b Banner) PrintMiddle() string {
	paddingsRep := strings.Repeat(" ", int(b.Padding))
	return fmt.Sprintf("%s%s%s%s%s", b.SidesChar, paddingsRep, b.Message, paddingsRep, b.SidesChar)
}

// PrintFullBanner Prints the complete banner
func (b Banner) PrintFullBanner() (string, error) {
	horLineTop, err := b.PrintHorizontalEdges(HorLineTypeTop)
	if err != nil {
		return "", err
	}

	horLineBottom, err := b.PrintHorizontalEdges(HorLineTypeBottom)
	if err != nil {
		return "", err
	}

	middle := b.PrintMiddle()

	return fmt.Sprintf("%s\n%s\n%s", horLineTop, middle, horLineBottom), nil
}
