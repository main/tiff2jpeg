# tiff2jpeg

A tiny Go library that converts TIFF images to JPEG.

It decodes TIFF via `golang.org/x/image/tiff` and re-encodes to JPEG at maximum quality (100), so the conversion is as lossless as JPEG allows. No CGo, no external tools.

## Install

```
go get github.com/main/tiff2jpeg
```

## Usage

```go
package main

import (
	"log"
	"os"

	"github.com/main/tiff2jpeg"
)

func main() {
	f, err := os.Open("photo.tif")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	jpegBytes, err := tiff2jpeg.TIFF2JPEG(f)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("photo.jpg", jpegBytes, 0o644); err != nil {
		log.Fatal(err)
	}
}
```

`TIFF2JPEG` accepts any `io.Reader` (file, HTTP response body, bytes buffer) and returns the encoded JPEG as `[]byte`.

## License

MIT — see [LICENSE](LICENSE).
