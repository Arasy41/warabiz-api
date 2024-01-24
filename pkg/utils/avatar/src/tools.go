package src

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand"
	"strings"
	"time"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

func GenerateBackground(color color.Color) *image.RGBA {
	bgImg := image.NewRGBA(image.Rect(0, 0, 200, 200))
	draw.Draw(bgImg, bgImg.Bounds(), &image.Uniform{color}, image.ZP, draw.Src)
	return bgImg
}

func AddLabel(bgImg *image.RGBA, label LabelConfiguration) (image.Image, error) {
	bounds := bgImg.Bounds()
	imgWidth := bounds.Dx()
	imgHeight := bounds.Dy()
	dc := gg.NewContextForRGBA(bgImg)
	dc.DrawImage(bgImg, 0, 0)
	if err := dc.LoadFontFace(label.Font, label.FontSize); err != nil {
		return nil, err
	}
	x := float64(imgWidth / 2)
	y := float64((imgHeight / 2) - label.YPosition)
	maxWidth := float64(imgWidth) - 60.0
	dc.SetColor(label.Color)
	dc.DrawStringWrapped(label.Text, x, y, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)
	return dc.Image(), nil
}

func GetInitialName(username string) string {
	if username == "" {
		return ""
	}
	if strings.Contains(username, " ") {
		parts := strings.Split(username, " ")
		if len(parts) < 2 {
			return username[0:1]
		} else {
			return parts[0][0:1] + parts[1][0:1]
		}
	} else if strings.Contains(username, ".") {
		if strings.Count(username, ".") == 1 {
			if !strings.Contains(username, ".com") {
				parts := strings.Split(username, ".")
				if len(parts) < 2 {
					return username[0:1]
				} else {
					return parts[0][0:1] + parts[1][0:1]
				}
			} else {
				return username[0:1]
			}
		} else if strings.Count(username, ".") > 1 {
			parts := strings.Split(username, ".")
			if len(parts) < 2 {
				return username[0:1]
			} else {
				return parts[0][0:1] + parts[1][0:1]
			}
		}
	} else if strings.Contains(username, "_") {
		parts := strings.Split(username, "_")
		if len(parts) < 2 {
			return username[0:1]
		} else {
			return parts[0][0:1] + parts[1][0:1]
		}
	} else {
		return username[0:1]
	}
	return username[0:1]
}

func ResizeImage(logo image.Image, size uint) (image.Image, error) {
	img := resize.Resize(size, size, logo, resize.Lanczos3)
	return img, nil
}

func PickRandomColor() color.Color {

	cPallete := []color.RGBA{}
	cPallete = append(cPallete,
		color.RGBA{235, 29, 54, 1},
		color.RGBA{255, 110, 49, 1},
		color.RGBA{255, 178, 0, 1},
		color.RGBA{139, 197, 65, 1},
		color.RGBA{34, 169, 224, 1},
		color.RGBA{60, 64, 72, 1},
		color.RGBA{113, 49, 221, 1},
	)

	rand.Seed(time.Now().UnixNano())
	return cPallete[rand.Intn(len(cPallete))]
}
