package main

import (
	"image"
)

func analyzeAlpha(img image.Image) (bool, bool) {
	skip := true
	hasAlphaPixel := false
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			_, _, _, a := img.At(x, y).RGBA()
			// XXX this feels like a hack:
			if a == 65535 || a == 0xff {
				skip = false
			} else {
				hasAlphaPixel = true
			}
			// XXX could return early if one non-alpha pixel and one alpha pixel has been found.
		}
	}
	return skip, hasAlphaPixel
}
