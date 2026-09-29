package img

// This provides MesloLGS Nerd Font Mono as the render font instead of the bundled Hack font face.

import (
	_ "embed"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

//go:embed fonts/Meslo-Regular.ttf
var mesloRegular []byte

//go:embed fonts/Meslo-Bold.ttf
var mesloBold []byte

//go:embed fonts/Meslo-Italic.ttf
var mesloItalic []byte

//go:embed fonts/Meslo-BoldItalic.ttf
var mesloBoldItalic []byte

func mesloFace(fontBytes []byte, opts *truetype.Options) font.Face {
	parsed, _ := truetype.Parse(fontBytes)
	return truetype.NewFace(parsed, opts)
}

// mesloFont satisfies the same interface as gonvenience/font's Hack provider, so
// output.go can use it as a drop-in replacement.
type mesloFont struct{}

var Meslo = mesloFont{}

func (f mesloFont) Regular(opts *truetype.Options) font.Face { return mesloFace(mesloRegular, opts) }
func (f mesloFont) Bold(opts *truetype.Options) font.Face    { return mesloFace(mesloBold, opts) }
func (f mesloFont) Italic(opts *truetype.Options) font.Face  { return mesloFace(mesloItalic, opts) }
func (f mesloFont) BoldItalic(opts *truetype.Options) font.Face {
	return mesloFace(mesloBoldItalic, opts)
}
