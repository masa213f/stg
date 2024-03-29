package resource

import (
	_ "embed"

	"github.com/golang/freetype/truetype"
	"github.com/masa213f/stg/pkg/util"
	"golang.org/x/image/font"
)

//go:embed files/font/arcade_n.ttf
var rawDataFont []byte

const (
	fontSize      = 16
	fontSizeSmall = 8
)

var (
	FontArcade      *util.Font
	FontArcadeSmall *util.Font
)

func loadFont(data []byte, size int) (*util.Font, error) {
	tt, err := truetype.Parse(data)
	if err != nil {
		return nil, err
	}

	const dpi = 72
	ret := util.NewFont(truetype.NewFace(tt, &truetype.Options{
		Size:    float64(size),
		DPI:     dpi,
		Hinting: font.HintingFull,
	}), size)

	return ret, nil
}

func init() {
	var err error
	FontArcade, err = loadFont(rawDataFont, fontSize)
	if err != nil {
		panic(err)
	}
	FontArcadeSmall, err = loadFont(rawDataFont, fontSizeSmall)
	if err != nil {
		panic(err)
	}
}
