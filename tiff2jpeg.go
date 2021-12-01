package tiff2jpeg

import (
	"bufio"
	"bytes"
	"golang.org/x/image/tiff"
	"image/jpeg"
	"io"
)

func TIFF2JPEG(r io.Reader) ([]byte, error) {
	img, err := tiff.Decode(r)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	if err := jpeg.Encode(w, img, &jpeg.Options{Quality: 100}); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
