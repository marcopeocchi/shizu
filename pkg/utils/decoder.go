package utils

import (
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/webp"
)

func DecodeImageFromBytes(b []byte) (image.Image, string, error) {
	return image.Decode(bytes.NewReader(b))
}
