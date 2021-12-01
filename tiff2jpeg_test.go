package tiff2jpeg

import (
	"os"
	"testing"
)

func TestTIFF2JPEG(t *testing.T) {
	file, err := os.Open("./IMG_9801-4.tif")
	if err != nil {
		t.Error(err)
	}

	b, err := TIFF2JPEG(file)
	if err != nil {
		t.Error(err)
	}

	if err := os.WriteFile("result.tiff", b, os.FileMode(0766)); err != nil {
		t.Error(err)
	}
}
