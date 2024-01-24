package file

import (
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"

	"github.com/nfnt/resize"
)

// ============================================================

func DownloadImage(url string) (image.Image, error) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		return nil, err
	}
	defer res.Body.Close()
	m, _, err := image.Decode(res.Body)
	if err != nil {
		return nil, err
	}
	return m, err
}

func ResizeImage(logo image.Image, size uint) (image.Image, error) {
	img := resize.Resize(size, size, logo, resize.Lanczos3)
	return img, nil
}

func CreateJPEGFile(img image.Image, imgName string) error {
	outFile, err := os.Create(imgName)
	if err != nil {
		return err
	}
	defer outFile.Close()
	jpeg.Encode(outFile, img, &jpeg.Options{Quality: 100})
	return nil
}

func CreatePNGFile(img image.Image, imgName string) error {
	outFile, err := os.Create(imgName)
	if err != nil {
		return err
	}
	defer outFile.Close()
	encoder := png.Encoder{CompressionLevel: png.DefaultCompression}
	err = encoder.Encode(outFile, img)
	return nil
}
