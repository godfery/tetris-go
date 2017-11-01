package assets

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
)

func ArcadeFontImage() image.Image {
	b := MustAsset("arcadefont.png")
	img, _, err := image.Decode(bytes.NewBuffer(b))
	if err != nil {
		panic(fmt.Sprintf("assets: image.Decode error: %v", err))
	}
	return img
}
