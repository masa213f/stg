package resource

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
)

var (
	//go:embed files/image/majo.png
	rawDataImagePlayer []byte
	//go:embed files/image/obake.png
	rawDataImageObake []byte
	//go:embed files/image/pipo-btleffect036.png
	rawDataImageEffectFire []byte
	//go:embed files/image/pipo-btleffect042.png
	rawDataImageEffectIce []byte
	//go:embed files/image/shot.png
	rawDataImageShot []byte
	//go:embed files/image/background.png
	rawDataImageBackground []byte
	//go:embed files/image/cloud1_128x48.png
	rawDataImageCloud1 []byte
	//go:embed files/image/cloud2_128x48.png
	rawDataImageCloud2 []byte
	//go:embed files/image/cloud3_96x48.png
	rawDataImageCloud3 []byte
)

const (
	// The size of sub-image.
	subImageWidth  = 32
	subImageHeight = 32
)

var (
	ImagePlayer     []*ebiten.Image
	ImageObake      []*ebiten.Image
	ImageEffectFire []*ebiten.Image
	ImageEffectIce  []*ebiten.Image
	ImageShot       *ebiten.Image
	ImageBackground *ebiten.Image
	ImageCloud      []*ebiten.Image
)

// loadImage crates a ebiten.Image from raw image data.
func loadImage(data []byte) (*ebiten.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	ret, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// loadImages craetes a slice of ebiten.Image from some raw image data.
func loadImages(dataSlice [][]byte) ([]*ebiten.Image, error) {
	ret := make([]*ebiten.Image, len(dataSlice))
	for i, data := range dataSlice {
		img, err := loadImage(data)
		if err != nil {
			return nil, err
		}
		ret[i] = img
	}
	return ret, nil
}

// loadSubImages creats a slice of ebiten.Image from one raw image data.
// This function cuts out areas specified in the 2nd argument.
// Specify the area with {X, Y, Width, Height}
func loadSubImages(data []byte, rects [][4]int) ([]*ebiten.Image, error) {
	img, err := loadImage(data)
	if err != nil {
		return nil, err
	}

	ret := make([]*ebiten.Image, len(rects))
	for i, r := range rects {
		sub, ok := img.SubImage(image.Rectangle{
			Min: image.Point{X: r[0], Y: r[1]},
			Max: image.Point{X: r[0] + r[2], Y: r[1] + r[3]},
		}).(*ebiten.Image)
		if !ok {
			return nil, fmt.Errorf("cannot get sub image: %d", i)
		}
		ret[i] = sub
	}

	return ret, nil
}

func init() {
	var err error
	ImagePlayer, err = loadSubImages(rawDataImagePlayer, [][4]int{
		{subImageWidth * 0, subImageHeight * 2, subImageWidth, subImageHeight},
		{subImageWidth * 1, subImageHeight * 2, subImageWidth, subImageHeight},
		{subImageWidth * 2, subImageHeight * 2, subImageWidth, subImageHeight},
		{subImageWidth * 1, subImageHeight * 2, subImageWidth, subImageHeight},
	})
	if err != nil {
		panic(err)
	}

	ImageObake, err = loadSubImages(rawDataImageObake, [][4]int{
		{subImageWidth * 0, subImageHeight, subImageWidth, subImageHeight},
		{subImageWidth * 1, subImageHeight, subImageWidth, subImageHeight},
		{subImageWidth * 2, subImageHeight, subImageWidth, subImageHeight},
		{subImageWidth * 1, subImageHeight, subImageWidth, subImageHeight},
	})
	if err != nil {
		panic(err)
	}

	ImageEffectFire, err = loadSubImages(rawDataImageEffectFire, [][4]int{
		{subImageWidth * 0, 0, subImageWidth, subImageHeight},
		{subImageWidth * 1, 0, subImageWidth, subImageHeight},
		{subImageWidth * 2, 0, subImageWidth, subImageHeight},
		{subImageWidth * 3, 0, subImageWidth, subImageHeight},
		{subImageWidth * 4, 0, subImageWidth, subImageHeight},
		{subImageWidth * 5, 0, subImageWidth, subImageHeight},
		{subImageWidth * 6, 0, subImageWidth, subImageHeight},
		{subImageWidth * 7, 0, subImageWidth, subImageHeight},
	})
	if err != nil {
		panic(err)
	}

	ImageEffectIce, err = loadSubImages(rawDataImageEffectIce, [][4]int{
		{subImageWidth * 0, 0, subImageWidth, subImageHeight},
		{subImageWidth * 1, 0, subImageWidth, subImageHeight},
		{subImageWidth * 2, 0, subImageWidth, subImageHeight},
		{subImageWidth * 3, 0, subImageWidth, subImageHeight},
		{subImageWidth * 4, 0, subImageWidth, subImageHeight},
		{subImageWidth * 5, 0, subImageWidth, subImageHeight},
		{subImageWidth * 6, 0, subImageWidth, subImageHeight},
		{subImageWidth * 7, 0, subImageWidth, subImageHeight},
	})
	if err != nil {
		panic(err)
	}

	ImageShot, err = loadImage(rawDataImageShot)
	if err != nil {
		panic(err)
	}

	ImageBackground, err = loadImage(rawDataImageBackground)
	if err != nil {
		panic(err)
	}

	ImageCloud, err = loadImages([][]byte{rawDataImageCloud1, rawDataImageCloud2, rawDataImageCloud3})
	if err != nil {
		panic(err)
	}
}
